package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/fanliao/go-promise"
	"io/ioutil"
	"k8s.io/client-go/third_party/forked/golang/template"
	"k8s.io/client-go/util/jsonpath"
	"net/http"
	"net/url"
	"regexp"
)

type (
	AuthRequestHandler struct {
		Client *http.Client
		Auth   *common.Auth
	}
)

func (h *AuthRequestHandler) RequestHandlerFunc(req *http.Request) error {

	if len(req.Header.Get("Authorization")) == 0 {

		// get realm info
		realmTask := func() (r interface{}, err error) {

			res, err := h.Client.Do(req)
			if err != nil {
				return res, err
			}
			defer func() {
				_ = res.Body.Close()
			}()

			v := res.Header.Get(http.CanonicalHeaderKey("WWW-Authenticate"))
			regex := `(?P<key>\w+)="(?P<token>[^"]+)`
			ret := map[string]string{}

			nameIdMap := map[string]int{}
			re := regexp.MustCompile(regex)
			for k, v := range re.SubexpNames() {
				nameIdMap[v] = k
			}
			for _, m := range re.FindAllStringSubmatch(v, -1) {
				ret[m[nameIdMap["key"]]] = m[nameIdMap["token"]]
			}

			return ret, nil
		}

		// get token
		authTask := func(data interface{}) (r interface{}, err error) {

			realmMap := data.(map[string]string)
			realmUrl, _ := url.Parse(realmMap["realm"])
			q := realmUrl.Query()
			if h.Auth != nil {
				q.Set("account", h.Auth.UserName)
			}
			q.Set("service", realmMap["service"])
			if realmMap["scope"] != "" {
				q.Set("scope", realmMap["scope"])
			} else {
				q.Set("offline_token", "true")
			}
			realmUrl.RawQuery = q.Encode()
			authReq, err := http.NewRequest(req.Method, realmUrl.String(), nil)
			if err != nil {
				return authReq, err
			}
			if h.Auth != nil {
				authReq.SetBasicAuth(h.Auth.UserName, h.Auth.PassWord)
			}
			res, err := h.Client.Do(authReq)
			if err != nil {
				return nil, err
			}

			if res.StatusCode/100 != 2 {
				return nil, errors.New(fmt.Sprintf("Authorization failed, status code: %d, result: %s. ", res.StatusCode, util.ReadWithDefault(res.Body, "{}")))
			}

			defer func() {
				_ = res.Body.Close()
			}()

			j := jsonpath.New("j")
			err = j.Parse(`{$.token}`)
			if err != nil {
				return nil, err
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}

			var o interface{}
			err = json.Unmarshal(body, &o)
			if err != nil {
				return nil, err
			}

			results, err := j.FindResults(o)
			if err != nil {
				return nil, err
			}

			for _, r := range results {
				for _, v := range r {
					token, ok := template.PrintableValue(v)
					if !ok {
						return nil, fmt.Errorf("can'h print type %s", v.Type())
					}
					return token, nil
				}
			}

			return nil, errors.New("get token failed")
		}

		f, ok := promise.Start(realmTask).Pipe(authTask)
		if !ok {
			return errors.New("promise error")
		}

		token, err := f.Get()
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	return nil

}
