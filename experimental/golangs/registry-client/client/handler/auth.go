package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fanliao/go-promise"
	"k8s.io/client-go/util/jsonpath"
	"net/http"
	"net/url"
	"regexp"
)

type (
	AuthRequestHandler struct {
		Client   *http.Client
		URL      string
		UserName string
		Password string
	}
)

func (t *AuthRequestHandler) Do() func(req *http.Request) error {
	return func(req *http.Request) error {

		if len(req.Header.Get("Authorization")) == 0 {

			realmTask := func() (r interface{}, err error) {

				// get realm info
				res, err := t.Client.Do(req)
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

			authTask := func(data interface{}) (r interface{}, err error) {

				// get wBuf
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
				authReq.SetBasicAuth(t.UserName, t.Password)
				res, err := t.Client.Do(authReq)
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

				wBuf := new(bytes.Buffer)
				rBuf := new(bytes.Buffer)
				_, _ = rBuf.ReadFrom(res.Body)
				var nodesData interface{}
				err = json.Unmarshal(rBuf.Bytes(), &nodesData)
				if err != nil {
					return nil, err
				}
				err = j.Execute(wBuf, nodesData)
				if err != nil {
					return nil, err
				}

				token := wBuf.String()

				return token, nil
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
