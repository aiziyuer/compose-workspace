使用指南
===


## 开发约定

``` bash

# 创建虚拟环境
virtualenv env

# 加载虚拟环境
source env/bin/activate

# 保存当前环境的类库
pip freeze > requirements.txt 

# 恢复类库
pip install -r requirements.txt 

# 测试ansible
ansible-playbook -i development 00-prepare.yml

```

### FAQ

- [kubernetes-sigs/kubespray](https://github.com/kubernetes-sigs/kubespray)
- [yum包安装](https://raymii.org/s/tutorials/Ansible_-_Only_if_on_specific_distribution_or_distribution_version.html)
- [官方文档](https://ansible-tran.readthedocs.io/en/latest/docs/playbooks_roles.html)
- [官方最佳实践](https://docs.ansible.com/ansible/latest/user_guide/playbooks_best_practices.html)