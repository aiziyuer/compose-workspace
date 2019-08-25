#!/bin/bash
# 
# Created by L.STONE <web.developer.network@gmail.com>
# -------------------------------------------------------------
# 自动创建 Docker TLS 证书
# -------------------------------------------------------------

# 以下是配置信息
# --[BEGIN]------------------------------

CODE="localhost"
IP="192.168.85.88"
PASSWORD="Changeme_123"
COUNTRY="CN"
STATE="Zhejiang"
CITY="Hangzhou"
ORGANIZATION="LC"
ORGANIZATIONAL_UNIT="Dev"
COMMON_NAME="$IP"
EMAIL="ziyu0123456789@qq.com"

# --[END]--

# Generate CA key
openssl genrsa -aes256 -passout "pass:$PASSWORD" -out "ca-key-$CODE.pem" 4096
# Generate CA
openssl req -new -x509 -days 365 -key "ca-key-$CODE.pem" -sha256 -out "ca-$CODE.pem" -passin "pass:$PASSWORD" -subj "/C=$COUNTRY/ST=$STATE/L=$CITY/O=$ORGANIZATION/OU=$ORGANIZATIONAL_UNIT/CN=$COMMON_NAME/emailAddress=$EMAIL"
# Generate Server key
openssl genrsa -out "server-key-$CODE.pem" 4096

# Generate Server Certs.
openssl req -subj "/CN=$COMMON_NAME" -sha256 -new -key "server-key-$CODE.pem" -out server.csr

echo "subjectAltName = IP:$IP,IP:127.0.0.1" >> extfile.cnf
echo "extendedKeyUsage = serverAuth" >> extfile.cnf

openssl x509 -req -days 365 -sha256 -in server.csr -passin "pass:$PASSWORD" -CA "ca-$CODE.pem" -CAkey "ca-key-$CODE.pem" -CAcreateserial -out "server-cert-$CODE.pem" -extfile extfile.cnf


# Generate Client Certs.
rm -f extfile.cnf

openssl genrsa -out "key-$CODE.pem" 4096
openssl req -subj '/CN=client' -new -key "key-$CODE.pem" -out client.csr
echo extendedKeyUsage = clientAuth >> extfile.cnf
openssl x509 -req -days 365 -sha256 -in client.csr -passin "pass:$PASSWORD" -CA "ca-$CODE.pem" -CAkey "ca-key-$CODE.pem" -CAcreateserial -out "cert-$CODE.pem" -extfile extfile.cnf

rm -vf client.csr server.csr

chmod -v 0400 "ca-key-$CODE.pem" "key-$CODE.pem" "server-key-$CODE.pem"
chmod -v 0444 "ca-$CODE.pem" "server-cert-$CODE.pem" "cert-$CODE.pem"

# 打包客户端证书
mkdir -p "tls-client-certs-$CODE"
cp -f "ca-$CODE.pem" "cert-$CODE.pem" "key-$CODE.pem" "tls-client-certs-$CODE/"
cd "tls-client-certs-$CODE"
tar zcf "tls-client-certs-$CODE.tar.gz" *
mv "tls-client-certs-$CODE.tar.gz" ../
cd ..
rm -rf "tls-client-certs-$CODE"

# 拷贝服务端证书
mkdir -p /etc/docker/certs.d
cp "ca-$CODE.pem" "server-cert-$CODE.pem" "server-key-$CODE.pem" /etc/docker/certs.d/

# /etc/docker/daemon.json
# {
#   "tlsverify": true,
#   "tlscacert": "/etc/docker/certs.d/ca.pem",
#   "tlscert": "/etc/docker/certs.d/server-cert.pem",
#   "tlskey": "/etc/docker/certs.d/server-key.pem",
#   "hosts": ["tcp://0.0.0.0:2376", "unix:///var/run/docker.sock"]
# }

echo " - 修改 /etc/docker/daemon.json 文件"
cat <<EOF
vi /etc/docker/daemon.json
{
  "tlsverify": true,
  "tlscacert": "/etc/docker/certs.d/ca-$CODE.pem",
  "tlscert": "/etc/docker/certs.d/server-cert-$CODE.pem",
  "tlskey": "/etc/docker/certs.d/server-key-$CODE.pem",
  "hosts": ["tcp://0.0.0.0:2376", "unix:///var/run/docker.sock"]
}
EOF

# 拷贝客户端证书文件
# cp -v {ca,cert,key}.pem ~/.docker

# 客户端远程连接
# docker -H 192.168.1.130:2376 --tlsverify --tlscacert ~/.docker/ca.pem --tlscert ~/.docker/cert.pem --tlskey ~/.docker/key.pem ps -a
echo "docker -H $IP:2376 --tlsverify --tlscacert ~/.docker/ca-$CODE.pem --tlscert ~/.docker/cert-$CODE.pem --tlskey ~/.docker/key-$CODE.pem ps -a"

# 客户端使用 cURL 连接
# curl --cacert ~/.docker/ca.pem --cert ~/.docker/cert.pem --key ~/.docker/key.pem https://192.168.1.130:2376/containers/json
echo "curl --cacert ~/.docker/ca-$CODE.pem --cert ~/.docker/cert-$CODE.pem --key ~/.docker/key-$CODE.pem https://$IP:2376/containers/json"

echo -e "\e[1;32mAll be done.\e[0m"