FROM centos:7

# 设置环境全英文避免出错
ENV LC_ALL=en_US.utf8 \
    LANG=en_US.UTF-8 \
    LC_CTYPE=en_US.UTF-8

# initial pass
RUN echo 'root:root' | chpasswd

# 覆盖源
COPY etc/yum.repos.d/*.repo /etc/yum.repos.d/

RUN \
    rpm --rebuilddb  \
 && yum install -y nodejs  
 
RUN \
    npm install -g kong-dashboard@v3.6.0
    
ENV \
    KONG_HOST=kong \
    KONG_PORT=8001 \
    KONG_DASHBOARD_PORT=8080

CMD ["sh", "-c", "kong-dashboard start --port ${KONG_DASHBOARD_PORT} --kong-url http://${KONG_HOST}:${KONG_PORT} "]
