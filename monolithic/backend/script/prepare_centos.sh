#!/usr/bin/env bash

####################################
## 更新软件源和软件
####################################

sudo yum update; sudo yum upgrade

####################################
## 安装工具软件
####################################

sudo yum install epel-release htop wget unzip -y

####################################
## 安装PM2
####################################

# 安装nodejs和npm
sudo yum install node npm -y

node -v
npm -v

# 安装pm2
npm install -g pm2
# 查看pm2的版本
pm2 --version
# tab补全
pm2 completion install
# 创建pm2开机启动脚本
pm2 startup
# 设置pm2的开机启动
sudo systemctl enable pm2-${USER}

####################################
## 安装Golang
####################################

latest_version=1.20.1

wget https://dl.google.com/go/go$latest_version.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzf go$latest_version.linux-amd64.tar.gz
rm -fr go$latest_version.linux-amd64.tar.gz

echo "export GOROOT=/usr/local/go" >> ~/.bashrc
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=~/go" >> ~/.bashrc
source ~/.bashrc

go version

####################################
## 安装Docker
####################################

sudo yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

sudo yum install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo yum install -y docker-compose

sudo systemctl start docker
