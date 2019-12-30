package handler

import (
	"net/http"
)

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
