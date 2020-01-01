package registry

import (
	"github.com/aiziyuer/registry/client/handler"
	"github.com/aiziyuer/registry/client/util"
	"net/http"
)

type Registry struct {
	URL     string
	Client  *http.Client
	Handler *handler.Facade
}

func NewClient(c *http.Client, url string, username string, password string) *Registry {

	return &Registry{
		URL:    url,
		Client: c,
		Handler: &handler.Facade{
			Client: c,
			Patterns: map[string]handler.Handler{
				".+": {
					Requests: map[string]func(*http.Request, *map[string]interface{}) error{
						"auth": (&handler.AuthRequestHandler{
							Client:   c,
							UserName: username,
							Password: password,
						}).F(),
					},
					Responses: map[string]func(req *http.Response) error{},
				},
			},
		},
	}
}

func (r *Registry) do(req *http.Request) (*http.Response, error) {

	context := &map[string]interface{}{
		"URL": r.URL,
	}

	return r.doWithContext(req, context)
}

func (r *Registry) doWithContext(req *http.Request, context *map[string]interface{}) (*http.Response, error) {
	return r.Handler.DoWithContext(req, context)
}

func (r *Registry) Ping() error {

	req, _ := http.NewRequest("GET", util.Url(r.URL, "/v2/"), nil)
	resp, err := r.do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	return err
}
