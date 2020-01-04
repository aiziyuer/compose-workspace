package registry

import (
	"github.com/aiziyuer/registry/client/auth"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
)

type (

	Endpoint struct {
		Schema string
		Host   string
	}

	Registry struct {
		Auth          *auth.BasicAuth
		Endpoint      *Endpoint
		Client        *http.Client
		HandlerFacade *handler.Facade
	}
)

func NewClient(c *http.Client, endpoint *Endpoint, auth *auth.BasicAuth) *Registry {

	return &Registry{
		Auth: auth,
		Endpoint: &Endpoint{
			Schema: endpoint.Schema,
			Host:   endpoint.Host,
		},
		Client: c,
		HandlerFacade: &handler.Facade{
			Client: c,
			Patterns: map[string]handler.Handler{
				".+": {
					Requests: map[string]func(*http.Request, *map[string]interface{}) error{
						"auth": (&handler.AuthRequestHandler{
							Client:   c,
							Auth: auth,
						}).RequestHandlerFunc(),
					},
					Responses: map[string]func(req *http.Response) error{},
				},
			},
		},
	}
}

func (r *Registry) do(req *http.Request) (*http.Response, error) {
	return r.HandlerFacade.Do(req)
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
