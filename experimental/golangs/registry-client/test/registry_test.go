package test

import (
	"crypto/tls"
	"fmt"
	"github.com/aiziyuer/registry/client/registry"
	"github.com/aiziyuer/registry/client/util"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"testing"
)

var client *registry.Registry

func init() {

	// 测试环境以.env文件为准, 生产环境改成godotenv.Load()
	err := godotenv.Overload()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client = registry.NewClient(&http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}, &registry.Endpoint{
		Schema: os.Getenv("REGISTRY_SCHEMA"),
		Host:   os.Getenv("REGISTRY_HOST"),
	}, &registry.BasicAuth{
		UserName: os.Getenv("REGISTRY_USERNAME"),
		PassWord: os.Getenv("REGISTRY_PASSWORD"),
	})
}

func TestClient(t *testing.T) {
	_ = client.Ping()
}

func TestTags(t *testing.T) {
	output, _ := client.Tags("aiziyuer/centos")
	fmt.Println(util.PrettyFormat(output))
}
