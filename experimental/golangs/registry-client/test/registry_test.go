package test

import (
	"crypto/tls"
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
	"os"
	"testing"
)

func TestClient(t *testing.T) {

	c := &http.Client{
		Transport: &http.Transport{
			Proxy:       http.ProxyFromEnvironment,
			DialContext: nil,
			DialTLS:     nil,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	router := handler.Router{
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
	}

	req, _ := http.NewRequest("GET", "https://registry-1.docker.io/v2/", nil)
	resp, err := router.Do(req)

	fmt.Println(resp, err)
}
