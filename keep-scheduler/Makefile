all: local

local: fmt vet
	GOOS=linux GOARCH=amd64 go build  -o=bin/keep-scheduler ./cmd/scheduler

build:  local
	docker build --no-cache . -t registry.cn-hangzhou.aliyuncs.com/keep_colocation/keep_scheduler:2.33

push:   build
	docker push registry.cn-hangzhou.aliyuncs.com/keep_colocation/keep_scheduler:2.33

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

clean: fmt vet
	sudo rm -f keep-scheduler