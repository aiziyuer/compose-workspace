package registry

import (
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
)

type (
	BasicAuth struct {
		UserName string
		PassWord string
	}

	Endpoint struct {
		Schema string
		Host   string
	}

	Registry struct {
		Auth     *BasicAuth
		Endpoint *Endpoint
		Client   *http.Client
		Handler  *handler.Facade
	}
)

func NewClient(c *http.Client, endpoint *Endpoint, auth *BasicAuth) *Registry {

	return &Registry{
		Auth: &BasicAuth{
			UserName: auth.UserName,
			PassWord: auth.PassWord,
		},
		Endpoint: &Endpoint{
			Schema: endpoint.Schema,
			Host:   endpoint.Host,
		},
		Client: c,
		Handler: &handler.Facade{
			Client: c,
			Patterns: map[string]handler.Handler{
				".+": {
					Requests: map[string]func(*http.Request, *map[string]interface{}) error{
						"auth": (&handler.AuthRequestHandler{
							Client:   c,
							UserName: auth.UserName,
							Password: auth.PassWord,
						}).RequestHandlerFunc(),
					},
					Responses: map[string]func(req *http.Response) error{},
				},
			},
		},
	}
}

func (r *Registry) do(req *http.Request) (*http.Response, error) {
	return r.Handler.Do(req)
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
	resp, _ := r.Handler.Do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	return err
}
