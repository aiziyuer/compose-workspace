package handler

import "net/http"

type RenderHandler struct {
}

func (h *RenderHandler) F() func(*http.Request, *map[string]interface{}) error {
	return func(request *http.Request, m *map[string]interface{}) error {

	}
}
