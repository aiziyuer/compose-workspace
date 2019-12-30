package handler

import (
	"net/http"
)

type Factory struct {
	Client     *http.Client
	PatternMap map[string]Handler
}

type Handler struct {
	Requests  map[string]func(req *http.Request) error
	Responses map[string]func(req *http.Response) error
}

func DefaultProxy() *Handler {
	return &Handler{
		Requests: map[string]func(req *http.Request) error{
			//"auth": nil,
		},
		Responses: map[string]func(req *http.Response) error{},
	}
}

// 通用执行函数
func (r *Factory) Do(req *http.Request) (*http.Response, error) {

	// 截取request获得真正的api
	var api = "/v2/"

	for _, handler := range r.PatternMap[api].Requests {
		err := handler(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	for _, handler := range r.PatternMap[api].Responses {
		err := handler(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
