package registry

import (
	"github.com/aiziyuer/registry/client/common"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
)

type (
	Endpoint struct {
		Schema string
		Host   string
	}

	Registry struct {
		Auth          *common.Auth
		Endpoint      *Endpoint
		Client        *http.Client
		HandlerFacade *handler.Facade
	}
)

func NewClient(c *http.Client, endpoint *Endpoint, auth *common.Auth) *Registry {

	return &Registry{
		Auth: auth,
		Endpoint: &Endpoint{
			Schema: endpoint.Schema,
			Host:   endpoint.Host,
		},
		Client: c,
		HandlerFacade: &handler.Facade{
			Client: c,
			PatternHandlerMap: map[string]handler.Handler{
				".+": {
					Requests: map[string]handler.RequestHandlerFunc{
						"common": (&handler.AuthRequestHandler{
							Client: c,
							Auth:   auth,
						}).RequestHandlerFunc,
					},
					Responses: map[string]handler.ResponseHandlerFunc{},
				},
			},
		},
	}
}

func (r *Registry) Ping() error {

	q, err := handler.NewApiRequest(handler.ApiRequestInput{
		"Schema": r.Endpoint.Schema,
		"Host":   r.Endpoint.Host,
		"Token":  "",
	}, `
	{
		"Method": "GET",
		"Path": "/v2/",
		"Schema": "{{ .Schema }}",
		"Host": "{{ .Host }}",
		"Header": {
			"Content-Type": "application/json; charset=utf-8",
			"Authorization": "{{ .Token }}",
		},
		"Body": "",
	}
`)
	if err != nil {
		return err
	}

	req, _ := q.Wrapper()
	resp, _ := r.HandlerFacade.Do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	return err
}
