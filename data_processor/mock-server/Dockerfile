#源镜像
FROM golang:1.17 as builder
#MAINTAINER ZhangLianjun "zhanglianjun@xiaohongshu.com"
#工作目录作目录
RUN go env -w CGO_ENABLED=0
ENV TZ Asia/Shanghai
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
WORKDIR /workspace
ADD mock-server.tar.gz .
#FROM registry.cn-hangzhou.aliyuncs.com/uestc-keep/distroless-static-nonroot:nonroot
CMD ["./run/main.sh"]




