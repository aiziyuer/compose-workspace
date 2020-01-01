package handler

import (
	"errors"
	"github.com/aiziyuer/registry/client/util"
	"net/http"
	"reflect"
	"regexp"
)

type (
	RequestHandler interface {
		F() func(*http.Request, *map[string]interface{}) error
		FV2() func(interface{}, *map[string]interface{}) error
	}
)

type Handler struct {
	Requests   map[string]func(req *http.Request, context *map[string]interface{}) error
	Responses  map[string]func(req *http.Response) error
	RequestsV2 map[string]func(input interface{}, context *map[string]interface{}) error
}

type Facade struct {
	Client   *http.Client
	Patterns map[string]Handler
}

func (r *Facade) DoV2(input *interface{}) (*http.Response, error) {
	return r.DoWithContextV2(input, &map[string]interface{}{})
}

func (r *Facade) DoWithContextV2(input interface{}, context *map[string]interface{}) (*http.Response, error) {

	if reflect.TypeOf(input).String() != "string" {
		return nil, errors.New("not support input(not string)")
	}

	requestStr := (input).(string)
	requestMap, err := util.JsonX2Map(requestStr)
	if err != nil {
		return nil, err
	}
	path := requestMap["Path"].(string)

	// 截取request获得真正的api进行处理函数的查找并执行
	for pattern, handler := range r.Patterns {

		p, _ := regexp.Compile(pattern)

		if p.MatchString(path) {

			for _, handler := range handler.RequestsV2 {
				err := handler(input, context)
				if err != nil {
					return nil, err
				}
			}

			requestJson, _ := util.TemplateRenderByPong2(requestStr, *context)
			req, _ := util.Json2Request(requestJson)

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

	return nil, errors.New("not match any pattern, current pattern: " + path)
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
