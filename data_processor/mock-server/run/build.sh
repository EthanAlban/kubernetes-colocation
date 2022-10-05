#!/usr/bin/env bash

#IMAGE_NAME=quicksilver_go_monitor_image
IMAGE_NAME=keep_mock_server
docker rmi $(docker images | grep 172.17.15.242/keep_colocation/${IMAGE_NAME} | awk '{print $3}') --force
cd ..
#bee pack -be GOOS=linux -be CGO_LDFLAGS="-static" -be CGO_ENABLED=1 -be GOARCH=amd64 -be CC=x86_64-linux-musl-gcc -be CXX=x86_64-linux-musl-g++
bee pack -be GOOS=linux -be GOARCH=amd64

tag='feat-keep-mock-server-'$(date '+%Y%m%d%H%M%S')
echo $tag
#docker build -f  DockerFile --platform=linux/amd64  --network=host  -t registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag .
docker build -f  DockerFile --platform=linux/amd64  --network=host  -t 172.17.15.242/keep_colocation/${IMAGE_NAME}:$tag .
echo 'Build complete.'
#docker push registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag
docker push 172.17.15.242/keep_colocation/${IMAGE_NAME}:$tag

echo "Start Clearing"
#rm -rf go-monitor-deploy/*
#docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
#docker rmi $(docker images | grep "none" | awk '{print $3}') --force
#docker rmi registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag

echo 'Build complete.'echo 'Build complete.'
#echo registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag
echo 172.17.15.242/keep_colocation/${IMAGE_NAME}:$tag
