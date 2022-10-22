# Build the manager binary
FROM golang:1.17 as builder
WORKDIR /keep_colocation
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

FROM registry.cn-hangzhou.aliyuncs.com/uestc-keep/distroless-static-nonroot:nonroot
WORKDIR /
COPY --from=builder /keep_colocation .
#USER 65532:65532

ENTRYPOINT ["/manager"]

# docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t 172.17.15.242/keep_colocation/node-simulator:latest .
# docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t keep_colocation/node-simulator:latest .

# docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t registry.cn-hangzhou.aliyuncs.com/keep_colocation/node-simulator:latest .
# docker push registry.cn-hangzhou.aliyuncs.com/keep_colocation/node-simulator:latest
# make deploy IMG=registry.cn-hangzhou.aliyuncs.com/keep_colocation/node-simulator