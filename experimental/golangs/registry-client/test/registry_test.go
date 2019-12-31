package test

import (
	"crypto/tls"
	"github.com/aiziyuer/registry/client/registry"
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

	client := registry.DefaultClient(c, "https://registry-1.docker.io", "aiziyuer", os.Getenv("REGISTRY_PASSWORD"))
	_ = client.Ping()
}
