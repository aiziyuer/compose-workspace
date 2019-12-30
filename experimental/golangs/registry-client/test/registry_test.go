package test

import (
	"crypto/tls"
	"fmt"
	"github.com/aiziyuer/registry/client/handler"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {

	factory := handler.Factory{
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
		PatternMap: map[string]handler.Handler{
			"/v2/": handler.Handler{
				Requests: map[string]func(req *http.Request) error{
					//"auth": nil,
				},
				Responses: map[string]func(req *http.Response) error{},
			},
		},
	}

	req, _ := http.NewRequest("GET", "registry-1.docker.io", nil)
	resp, err := factory.Do(req)

	fmt.Println(resp, err)
}
