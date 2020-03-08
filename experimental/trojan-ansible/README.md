trojan auto apply
===

archive a trojan ploybook.

### prepare

``` bash

yum install -y \
    gcc python36-devel openssl-devel sshpass openssh-clients

pip3 install -r requirementes.txt

# prevent from ssh failed
cat <<'EOF'>/root/.ssh/config
Host *
    StrictHostKeyChecking no
    UserKnownHostsFile=/dev/null
EOF

```

### how to use

``` bash

# fork a config from example
cp inventory/example.ini inventory/{domain}.ini 

# apply
ansible-playbook install.yaml -i inventory/{domain}.ini

```


