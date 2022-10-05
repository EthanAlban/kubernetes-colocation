#!/usr/bin/env bash

IMAGE_NAME=node-simulator
docker rmi $(docker images | grep 172.17.15.242/keep_colocation/${IMAGE_NAME} | awk '{print $3}') --force
cd ..

tag='feat-keep-node-simulator-'$(date '+%Y%m%d%H%M%S')
echo $tag
#docker build -f  DockerFile --platform=linux/amd64  --network=host  -t registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag .
docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t 172.17.15.242/keep_colocation/node-simulator:latest .
echo 'Build complete.'
docker push 172.17.15.242/keep_colocation/${IMAGE_NAME}:$tag

echo "Start Clearing"
#rm -rf go-monitor-deploy/*
docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker rmi $(docker images | grep "none" | awk '{print $3}') --force
#docker rmi registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag

echo 'Build complete.'echo 'Build complete.'
#echo registry.cn-hangzhou.aliyuncs.com/uyistcoj/go-monitor.mlsys.xiaohongshu.com:$tag
echo 172.17.15.242/keep_colocation/${IMAGE_NAME}:$tag
