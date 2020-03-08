trojan自动安装
===

归档一下自动初始化一套trojan系统的编排任务

### ansible的运行环境

``` bash

# 当前只支持centos7
yum install -y \
    gcc python36-devel openssl-devel sshpass openssh-clients

# 恢复类库
pip3 install -r requirementes.txt

# 清除所有虚环境的包 !! 危险 !!
#pip3 freeze | xargs pip3 uninstall -y

# ssh_config设置样例
cat <<'EOF'>/root/.ssh/config
Host *
    StrictHostKeyChecking no
    UserKnownHostsFile=/dev/null
EOF

```

### 部署任务

``` bash

ansible-playbook -i inventory/example.ini trojan-install.yaml

```


