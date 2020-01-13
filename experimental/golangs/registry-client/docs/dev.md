### 命令行生成

``` bash

# 下载命令行脚手架代码生成器
go get github.com/spf13/cobra/cobra

# 生成主命令
cobra add -l none registryV2 --pkg-name main

# 生成子命令
cobra add image

```