
GOPATH:=$(shell go env GOPATH)


.PHONY: proto build test
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/vessel/vessel.proto

build: proto
	GOOS=linux GOARCH=amd64 go build -o vessel-srv main.go
	docker build -t vessel-srv .

build-deploy: proto
	GOOS=linux GOARCH=amd64 go build -o vessel-srv main.go
	docker build -t gcr.io/impactful-yeti-230307/vessel-srv:latest .

test:
	go test -v ./... -cover