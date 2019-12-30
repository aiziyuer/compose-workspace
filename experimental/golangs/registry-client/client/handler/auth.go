package handler

import "net/http"

type (
	AuthRequestHandler struct {
		URL      string
		UserName string
		Password string
	}
)

func (t *AuthRequestHandler) Do() func(req *http.Request) error {
	return func(req *http.Request) error {

		return nil
	}
}
