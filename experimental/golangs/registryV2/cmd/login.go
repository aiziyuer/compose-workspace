package cmd

import (
	"crypto/tls"
	"encoding/base64"
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/registry"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type (
	ConfigStore struct {
		Auths map[string]map[string]string `json:"auths"`
	}
)

func init() {

	var name, pass string

	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login docker registry",
		RunE: func(cmd *cobra.Command, args []string) error {

			c := registry.NewClient(&http.Client{
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
				UserName: name,
				PassWord: pass,
			})

			if err := c.Login(); err != nil {
				return err
			}

			home, err := homedir.Dir()
			if err != nil {
				return err
			}

			confDir := path.Join(home, ".registryV2")
			if err := os.MkdirAll(confDir, os.ModePerm); err != nil {
				return err
			}

			// 读取存储密码
			config := &ConfigStore{
				Auths: map[string]map[string]string{},
			}
			configFile := path.Join(confDir, "config.json")
			if _, err := os.Stat(configFile); err == nil {
				json, err := ioutil.ReadFile(configFile)
				if err != nil {
					return err
				}

				if err := util.JsonX2Object(string(json), config); err != nil {
					return err
				}
			}

			// 更新密码
			resultMap := config.Auths[c.Endpoint.Host]
			if resultMap != nil {
				resultMap["auth"] = base64.StdEncoding.EncodeToString([]byte(name + ":" + pass))
			} else {
				config.Auths[c.Endpoint.Host] = map[string]string{
					"auth": base64.StdEncoding.EncodeToString([]byte(name + ":" + pass)),
				}
			}

			// 写入存储密码
			jsonBytes, err := util.Object2JsonBytes(config)
			if err != nil {
				return err
			}

			jsonBytes = util.PrettyJsonBytes(jsonBytes)
			if err := ioutil.WriteFile(configFile, jsonBytes, os.ModePerm); err != nil {
				return err
			}

			return nil
		},
	}
	rootCmd.AddCommand(loginCmd)

	loginCmd.PersistentFlags().StringVarP(&name, "name", "u", "", "user name")
	loginCmd.PersistentFlags().StringVarP(&pass, "pass", "p", "", "user password")
	_ = loginCmd.MarkPersistentFlagRequired("name")
	_ = loginCmd.MarkPersistentFlagRequired("pass")
}
