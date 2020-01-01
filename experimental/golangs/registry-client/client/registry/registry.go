package registry

import (
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
	"os"
	"regexp"
)

type Registry struct {
	URL    string
	Client *http.Client
	Router *handler.Facade
}

func DefaultClient(c *http.Client, url string, username string, password string) *Registry {

	return &Registry{
		URL:    url,
		Client: c,
		Router: &handler.Facade{
			Client: c,
			Patterns: map[string]handler.Handler{
				".+": {
					Requests: map[string]func(req *http.Request) error{
						"auth": (&handler.AuthRequestHandler{
							Client:   c,
							UserName: "aiziyuer",
							Password: os.Getenv("REGISTRY_PASSWORD"),
						}).Do(),
					},
					Responses: map[string]func(req *http.Response) error{},
				},
			},
		},
	}
}

func (r *Registry) do(req *http.Request) (*http.Response, error) {
	return r.Router.Do(req)
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
