package test

import (
	"crypto/tls"
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {

	router := handler.Router{
		Client: &http.Client{
			Transport: &http.Transport{
				Proxy:       http.ProxyFromEnvironment,
				DialContext: nil,
				DialTLS:     nil,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Patterns: map[string]handler.Handler{
			".+": {
				Requests: map[string]func(req *http.Request) error{
					"auth": (&handler.AuthRequestHandler{
						UserName: "aiziyuer",
						Password: "aiziyuer",
					}).Do(),
				},
				Responses: map[string]func(req *http.Response) error{},
			},
		},
	}

	req, _ := http.NewRequest("GET", "registry-1.docker.io", nil)
	resp, err := router.Do(req)

	fmt.Println(resp, err)
}
