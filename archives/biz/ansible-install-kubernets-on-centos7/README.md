功能说明
===

ansible 安装kubernetes


### 测试功能

```
# 启动应用
docker-compose build && docker-compose up -d 

# 进入ansible工作机
docker-compose exec ansible bash

# ansible安装
ansible-playbook site.yml

```


### FAQ

- [ansible最佳实践](https://docs.ansible.com/ansible/latest/user_guide/playbooks_best_practices.html)