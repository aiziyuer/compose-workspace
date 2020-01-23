package cmd

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/registry"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"path"
	"strings"
)

var isDebug bool

var rootCmd = &cobra.Command{
	Use: "registryV2",
	//TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if isDebug {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.SetReportCaller(true)
		}
	},
}

var outputFormat string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {

	// 默认关闭调试开关
	isDebug = false

	rootCmd.PersistentFlags().StringVarP(
		&outputFormat,
		"output", "o", "table",
		"options output format: table, yaml, json ",
	)

	rootCmd.PersistentFlags().BoolVarP(
		&isDebug,
		"debug", "d", true,
		"show verbose log ",
	)

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
