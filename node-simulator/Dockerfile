# Build the manager binary
FROM golang:1.17 as builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

FROM registry.cn-hangzhou.aliyuncs.com/uestc-keep/distroless-static-nonroot:nonroot
WORKDIR /
COPY --from=builder /workspace .
#USER 65532:65532

ENTRYPOINT ["/manager"]

# docker build -f  Dockerfile --platform=linux/amd64  --network=host  -t 172.17.15.242/keep_colocation/node-simulator:latest .
# docker push 172.17.15.242/keep_colocation/node-simulator:latest