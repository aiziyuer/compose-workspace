package handler

import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/aiziyuer/registryV2/impl/util"
	"github.com/sethgrid/pester"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	textTemplate "text/template"
)

type (
	RequestHandlerFunc  func(req *http.Request) error
	ResponseHandlerFunc func(resp *http.Response) error

	Handler struct {
		RequestFns  map[string]RequestHandlerFunc
		ResponseFns map[string]ResponseHandlerFunc
	}

	Facade struct {
		Client *http.Client
		// url => handlers
		PatternHandlerMap map[string]Handler
	}

	ApiRequestInput map[string]interface{}
	ApiRequest      struct {
		Input    map[string]interface{}
		Template string
		Output   string
		Method   string
		Path     string
		Schema   string
		Host     string
		URL      string
		Params   map[string]string
		Headers  map[string]interface{}
		Body     interface{}
	}
)

func NewApiRequest(template string, input map[string]interface{}) (*ApiRequest, error) {
	return (&ApiRequest{
		Input:    input,
		Template: template,
	}).Render()
}

func (r *ApiRequest) Render() (*ApiRequest, error) {

	var output bytes.Buffer
	t, err := textTemplate.New("").Parse(r.Template)
	if err != nil {
		return nil, err
	}
	if err := t.Execute(&output, r.Input); err != nil {
		logrus.Debug("textTemplate render failed")
		logrus.Debugf("template: %s", r.Template)
		logrus.Debugf("input: %s", r.Input)
		return nil, err
	}

	r.Output = output.String()
	if err := JsonX2Object(r.Output, &r); err != nil {
		return nil, err
	}

	r.URL = fmt.Sprintf("%s://%s%s", r.Schema, r.Host, r.Path)

	return r, nil
}

func (r *ApiRequest) Wrapper() (*http.Request, error) {

	var body io.Reader
	body = nil
	switch t := r.Body.(type) {
	case string:
		if len(t) > 0 {
			body = strings.NewReader(r.Body.(string))
		}
	case map[string]string:
		if len(t) > 0 {
			values := &url.Values{}
			for k, v := range r.Body.(map[string]string) {
				values.Set(k, v)
			}
			body = strings.NewReader(values.Encode())
		}
	default:
		body = nil
	}

	req, err := http.NewRequest(r.Method, r.URL, body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		if k != "" && v != "" {
			header := req.Header.Get(k)
			if header == "" {
				switch v.(type) {
				case string:
					x := v.(string)
					req.Header.Set(k, x)
				case []interface{}:
					x := v.([]interface{})
					if len(x) >= 1 {
						req.Header.Set(k, x[0].(string))
						for i := 1; i < len(x); i++ {
							req.Header.Add(k, x[i].(string))
						}
					}
				}

			} else {
				switch v.(type) {
				case string:
					x := v.(string)
					req.Header.Add(k, x)
				case []interface{}:
					x := v.([]interface{})
					for i := 0; i < len(x); i++ {
						req.Header.Add(k, x[i].(string))
					}

				}

			}
		}
	}

	for k, v := range r.Params {
		q := req.URL.Query()
		if k != "" && v != "" {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return req, nil
}

// 通用执行函数
func (r *Facade) Do(req *http.Request) (*http.Response, error) {

	// 截取request获得真正的api进行处理函数的查找并执行
	for pattern, handler := range r.PatternHandlerMap {

		p, _ := regexp.Compile(pattern)
		if p.MatchString(req.URL.Path) {

			for _, fn := range handler.RequestFns {
				err := fn(req)
				if err != nil {
					return nil, err
				}
			}

			//resp, err := r.Client.Do(req)
			c := pester.NewExtendedClient(r.Client)
			c.MaxRetries = 3
			resp, err := c.Do(req)

			if err != nil {
				return nil, err
			}

			for _, fn := range handler.ResponseFns {
				err := fn(resp)
				if err != nil {
					return nil, err
				}
			}

			return resp, nil
		}
	}

	return nil, errors.New("not match any pattern, current pattern: " + req.URL.Path)
}
