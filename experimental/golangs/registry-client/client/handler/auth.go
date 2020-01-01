package handler

import (
	"encoding/json"
	"errors"
	"fmt"
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
		Client   *http.Client
		UserName string
		Password string
	}
)

func (h *AuthRequestHandler) FV2() func(interface{}, *map[string]interface{}) error {
	return func(i interface{}, m *map[string]interface{}) error {

		return nil
	}
}

func (h *AuthRequestHandler) F() func(*http.Request, *map[string]interface{}) error {
	return func(req *http.Request, context *map[string]interface{}) error {

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
				q.Set("offline_token", "true")
				q.Set("service", realmMap["service"])
				if realmMap["scope"] != "" {
					q.Set("scope", realmMap["scope"])
				}
				realmUrl.RawQuery = q.Encode()
				authReq, err := http.NewRequest(req.Method, realmUrl.String(), nil)
				if err != nil {
					return authReq, err
				}
				authReq.SetBasicAuth(h.UserName, h.Password)
				res, err := h.Client.Do(authReq)
				if err != nil {
					return nil, err
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
}
