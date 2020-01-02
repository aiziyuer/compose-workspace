package handler

import (
	"bytes"
	"errors"
	"fmt"
	. "github.com/aiziyuer/registry/client/util"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	textTemplate "text/template"
)

type (
	RequestHandler interface {
		RequestHandlerFunc() func(*http.Request, *map[string]interface{}) error
	}

	Handler struct {
		Requests  map[string]func(req *http.Request, context *map[string]interface{}) error
		Responses map[string]func(req *http.Response) error
	}

	Facade struct {
		Client   *http.Client
		Patterns map[string]Handler
	}

	ApiRequestInput map[string]interface{}
	apiRequest      struct {
		Input    map[string]interface{}
		Template string
		Output   string
		Method   string
		Path     string
		Schema   string
		Host     string
		URL      string
		Params   map[string]string
		Headers  map[string]string
		Body     interface{}
	}
)

func NewApiRequest(input map[string]interface{}, template string) (*apiRequest, error) {
	return (&apiRequest{
		Input:    input,
		Template: template,
	}).Render()
}

func (r *apiRequest) Render() (*apiRequest, error) {

	var output bytes.Buffer
	t, err := textTemplate.New("").Parse(r.Template)
	if err != nil {
		return nil, err
	}
	if err := t.Execute(&output, r.Input); err != nil {
		return nil, err
	}

	r.Output = output.String()
	if err := JsonX2Object(r.Output, &r); err != nil {
		return nil, err
	}

	r.URL = fmt.Sprintf("%s://%s%s", r.Schema, r.Host, r.Path)

	return r, nil
}

func (r *apiRequest) Wrapper() (*http.Request, error) {

	var body io.Reader
	switch r.Body.(type) {
	case string:
		body = strings.NewReader(r.Body.(string))
	case map[string]string:
		values := &url.Values{}
		for k, v := range r.Body.(map[string]string) {
			values.Set(k, v)
		}
		body = strings.NewReader(values.Encode())
	default:
		body = nil
	}

	req, err := http.NewRequest(r.Method, r.URL, body)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers {
		if k != "" && v != "" {
			req.Header.Set(k, v)
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

func (r *Facade) Do(req *http.Request) (*http.Response, error) {
	return r.DoWithContext(req, &map[string]interface{}{})
}

// 通用执行函数
func (r *Facade) DoWithContext(req *http.Request, context *map[string]interface{}) (*http.Response, error) {

	// 截取request获得真正的api进行处理函数的查找并执行
	for pattern, handler := range r.Patterns {

		p, _ := regexp.Compile(pattern)
		if p.MatchString(req.URL.Path) {

			for _, handler := range handler.Requests {
				err := handler(req, context)
				if err != nil {
					return nil, err
				}
			}

			resp, err := r.Client.Do(req)
			if err != nil {
				return nil, err
			}

			for _, handler := range handler.Responses {
				err := handler(resp)
				if err != nil {
					return nil, err
				}
			}

			return resp, nil
		}

	}

	return nil, errors.New("not match any pattern, current pattern: " + req.URL.Path)
}
