#!/usr/bin/env bash

IMAGE_NAME=node-simulator
docker rmi $(docker images | grep  registry.cn-hangzhou.aliyuncs.com/keep_colocation/${IMAGE_NAME} | awk '{print $3}') --force
cd ..

echo "vendoring current project"
go mod vendor
tag='feat-keep-node-simulator-'$(date '+%Y%m%d%H%M%S')
echo $tag
docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t  registry.cn-hangzhou.aliyuncs.com/keep_colocation/node-simulator:latest .
echo 'Build complete.'
docker push registry.cn-hangzhou.aliyuncs.com/keep_colocation/${IMAGE_NAME}:latest

echo "Start Clearing"
docker rm $(docker ps -a | grep "Exited" | awk '{print $1 }')
docker rmi $(docker images | grep "none" | awk '{print $3}') --force

echo 'Clear complete.'
echo  registry.cn-hangzhou.aliyuncs.com/keep_colocation/${IMAGE_NAME}:latest
echo "rm -rf vendor"
rm -rf vendor