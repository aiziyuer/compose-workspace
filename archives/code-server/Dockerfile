FROM centos:7

# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

RUN \
    yum install -y deltarpm \
    && yum clean all

# 开发套件
RUN \
    yum group install -y 'Development Tools' \
    && yum clean all

# 设置第三方源
RUN \
    yum install -y epel-release \
    yum install -y https://repo.ius.io/ius-release-el7.rpm

# 常用的工具
RUN \
    yum install -y \
        vim net-tools git224 \
        tcpdump openssh-clients \
        zsh zip unzip telnet \
        lsof wget curl \
    && yum clean all

# 容器相关的工具
RUN \
    yum install -y yum-utils \
    && yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo \
    && yum install -y docker-ce docker-ce-cli containerd.io \
    && yum install -y docker-compose \
    && yum clean all

RUN \
    mv /root /root-original \
    && ln -s root-original root

# oh-my-zsh
RUN \
    sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# 各类SDK
# GVM
RUN \
    curl -s -S -L -o- https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash \
    && yum install -y golang \
    && yum clean all \
    && echo '[[ -s "/root/.gvm/scripts/gvm" ]] && source "/root/.gvm/scripts/gvm"' >> ~/.zshrc \
    && . ~/.bashrc && gvm install go1.4 -B \
    && gvm use go1.4 && export GOROOT_BOOTSTRAP=$GOROOT \
    && gvm install go1.5 \
    && gvm install go1.16 \
    && gvm use go1.16 --default

# NVM
RUN \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.37.2/install.sh | bash \
    && echo 'export NVM_DIR="$HOME/.nvm"' >> ~/.zshrc \
    && echo '[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm' >> ~/.zshrc \
    && echo '[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion' >> ~/.zshrc \
    && . ~/.bashrc && nvm install --lts && nvm use --lts --default

# SDKMan
RUN \
    curl -s "https://get.sdkman.io" | bash \
    && echo '#THIS MUST BE AT THE END OF THE FILE FOR SDKMAN TO WORK!!!' >> ~/.zshrc \
    && echo 'export SDKMAN_DIR="/root/.sdkman"' >> ~/.zshrc \
    && echo '[[ -s "/root/.sdkman/bin/sdkman-init.sh" ]] && source "/root/.sdkman/bin/sdkman-init.sh"' >> ~/.zshrc \
    && . ~/.bashrc && sdk install maven && sdk install gradle && sdk install java

# 升级code-server
RUN \
    yum install -y https://github.com/cdr/code-server/releases/download/v3.9.0/code-server-3.9.0-amd64.rpm

RUN \
    (cd /lib/systemd/system/sysinit.target.wants/; for i in *; do [ $i == systemd-tmpfiles-setup.service ] || rm -f $i; done); \
    rm -f /lib/systemd/system/multi-user.target.wants/*;\
    rm -f /etc/systemd/system/*.wants/*;\
    rm -f /lib/systemd/system/local-fs.target.wants/*; \
    rm -f /lib/systemd/system/sockets.target.wants/*udev*; \
    rm -f /lib/systemd/system/sockets.target.wants/*initctl*; \
    rm -f /lib/systemd/system/basic.target.wants/*;\
    rm -f /lib/systemd/system/anaconda.target.wants/*;

RUN \
    systemctl enable code-server@root.service 

# 避免出现僵尸进程: https://developers.redhat.com/blog/2016/09/13/running-systemd-in-a-non-privileged-container/
STOPSIGNAL SIGRTMIN+3

CMD ["/usr/sbin/init"]

LABEL "description"="usage: rsync -av /root-original/ /root/"
