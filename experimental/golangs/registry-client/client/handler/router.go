package handler

import "net/http"

type Router struct {
	Client   *http.Client
	Patterns map[string]Handler
}

// 通用执行函数
func (r *Router) Do(req *http.Request) (*http.Response, error) {

	// 截取request获得真正的api
	var api = "/v2/"

	for _, handler := range r.Patterns[api].Requests {
		err := handler(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, err
	}

	for _, handler := range r.Patterns[api].Responses {
		err := handler(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
