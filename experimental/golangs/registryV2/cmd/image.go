package cmd

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/registry"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"net/http"
	"path"
	"strings"
)

func init() {
	var imageCmd = &cobra.Command{
		Use:   "image",
		Short: "Image search/inspect ",
	}
	rootCmd.AddCommand(imageCmd)

	var imageSearchCmd = &cobra.Command{
		Use:   "search",
		Short: "Search by name",
		RunE: func(cmd *cobra.Command, args []string) error {

			c, err2 := getClient()
			if err2 != nil {
				return err2
			}

			_ = c.Ping()

			// TODO 业务逻辑

			return errors.New("")
		},
	}
	imageCmd.AddCommand(imageSearchCmd)

}

func getClient() (*registry.Registry, error) {

	var client *registry.Registry = nil

	home, err := homedir.Dir()
	if err != nil {
		return client, err
	}

	// 尝试读取存储密码
	jsonParsed, _ := gabs.ParseJSONFile(path.Join(home, ".registryV2/config.json"))
	encodedAuth := jsonParsed.
		Search("auths").
		Search(util.GetEnvAnyWithDefault("registry-1.docker.io", "REGISTRY_HOST")).
		Search("auth").
		Data()

	if encodedAuth != nil {
		ret, err := base64.StdEncoding.DecodeString(encodedAuth.(string))
		if err == nil {
			nameAndPass := string(ret)

			client = registry.NewClient(&http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}, &registry.Endpoint{
				Schema: util.GetEnvAnyWithDefault("https", "REGISTRY_SCHEMA"),
				Host:   util.GetEnvAnyWithDefault("registry-1.docker.io", "REGISTRY_HOST"),
			}, &common.Auth{
				UserName: strings.Split(nameAndPass, ":")[0],
				PassWord: strings.Split(nameAndPass, ":")[1],
			})
		}

	}

	// 实在不行就采用无认证客户端
	if client == nil {
		client = registry.NewClient(&http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}, &registry.Endpoint{
			Schema: util.GetEnvAnyWithDefault("https", "REGISTRY_SCHEMA"),
			Host:   util.GetEnvAnyWithDefault("registry-1.docker.io", "REGISTRY_HOST"),
		}, nil)

	}

	return client, nil
}
