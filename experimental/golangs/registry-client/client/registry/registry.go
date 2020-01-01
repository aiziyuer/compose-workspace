package registry

import (
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
	"regexp"
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
					Requests: map[string]func(*http.Request, *map[string]string) error{
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

	context := &map[string]string{
		"URL": r.URL,
	}

	return r.doWithContext(req, context)
}

func (r *Registry) doWithContext(req *http.Request, context *map[string]string) (*http.Response, error) {
	return r.Handler.Do(req, context)
}

func (r *Registry) url(pathTemplate string, args ...interface{}) string {

	tmpUrl := fmt.Sprintf("%s/%s", r.URL, fmt.Sprintf(pathTemplate, args...))

	m := r.regexNamedMatch(tmpUrl, `(?P<schema>^\w+://|^)(?P<host>[^/]+)(?P<path>[\w/]+)`)
	schema := m["schema"]
	if schema == "" {
		schema = "https://"
	}
	host := m["host"]
	path := m["path"]

	tmpSuffix := fmt.Sprintf("%s/%s", host, path)
	suffix := regexp.MustCompile(`[/]+`).ReplaceAllString(tmpSuffix, `/`)
	url := fmt.Sprintf("%s%s", schema, suffix)
	return url
}

func (r *Registry) Ping() error {

	req, _ := http.NewRequest("GET", r.url("/v2/"), nil)
	resp, err := r.do(req)
	if resp != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
	}
	return err
}

func (r *Registry) regexNamedMatch(input string, pattern string) map[string]string {

	re := regexp.MustCompile(pattern)
	m := map[string]string{}
	for i, v := range re.FindStringSubmatch(input) {
		name := re.SubexpNames()[i]
		if name == "" {
			continue
		}
		m[name] = v
	}

	return m
}
