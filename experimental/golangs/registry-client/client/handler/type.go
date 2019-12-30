package handler

import "net/http"

type (
	RequestHandler interface {
		Do() func(req *http.Request) error
	}
)
