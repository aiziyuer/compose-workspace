
package main

import (
	"github.com/huaweicse/iam"
	"fmt"
)

func main() {
	
	var c *iam.Client
	var err error
	c, err = iam.NewClient(iam.Options{
		Endpoint:  "https://iam.cn-north-1.myhuaweicloud.com",
		AccessKey: "xxx",
		SecretKey: "xxx",
		Project:   "default",
	})

	fmt.Println("hello world", c, err)
}
