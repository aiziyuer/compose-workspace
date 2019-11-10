项目说明
===

这里主要归档平时对kubernetes的自动化安装过程

### 协作约定

``` bash
yum install -y gcc python36-devel openssl-devel
yum install -y sshpass openssh-clients
pip3 install virtualenv

# 创建虚拟环境
virtualenv .env

# 加载虚环境
source .env/bin/activate

# 保存当前类库
# pip3 freeze > requirementes.txt

# 恢复类库
pip3 install -r requirementes.txt

# 清除所有虚环境的包 !! 危险 !!
pip3 freeze | xargs pip3 uninstall -y

```


