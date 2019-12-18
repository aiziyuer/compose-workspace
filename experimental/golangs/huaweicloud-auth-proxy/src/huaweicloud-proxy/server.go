
package main

import (
	"github.com/gophercloud/gophercloud/auth/aksk"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack"
	"fmt"
)


func main() {
	//设置认证参数
	akskOpts := aksk.AKSKOptions{
		IdentityEndpoint: "https://iam.example.com/v3",
		DomainID:         "{domainid}",
		ProjectID:        "{projectid}",
		Cloud:            "myhuaweicloud.com",
		Region:           "cn-north-1",
		AccessKey:        "{your AK string}",
		SecretKey:        "{your SK string}",
	}
	//初始化provider client
	provider, providerErr := openstack.AuthenticatedClient(akskOpts)
	if providerErr != nil {
		fmt.Println("init provider client error:", providerErr)
		panic(providerErr)
	}

	//初始化service client
	sc, serviceErr := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if serviceErr != nil {
		fmt.Println("init compute service client error:", serviceErr)
		panic(serviceErr)
	}

	//列出所有服务器
	allPages, err := servers.List(sc, servers.ListOpts{}).AllPages()

	if err != nil {
		fmt.Println("request server list error:", err)
		panic(err)
	}
	//解析返回值
	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		fmt.Println("extract response data error:", err)
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	//打印信息
	fmt.Println("List Servers:")
	for _, s := range allServers {
		fmt.Println("server ID is :", s.ID)
		fmt.Println("server name is :", s.Name)
		fmt.Println("server Status is :", s.Status)
		fmt.Println("server AvailbiltyZone is :", s.AvailbiltyZone)
	}
}