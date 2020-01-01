package test

import (
	"crypto/tls"
	"github.com/aiziyuer/registry/client/registry"
	"net/http"
	"os"
	"testing"
)

var client *registry.Registry

func init() {

	client = registry.NewClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}, "https://registry-1.docker.io", "aiziyuer", os.Getenv("REGISTRY_PASSWORD"))
}

func TestClient(t *testing.T) {
	_ = client.Ping()
}

func TestTags(t *testing.T) {
	_, _ = client.Tags("aiziyuer/centos")
}
