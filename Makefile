ITERATION := $(shell date +%s)
ENV ?= "$(shell uname -n)"
COMMIT ?= $$(git rev-parse HEAD)
TAG ?= $$(git describe --tags --abbrev=0 2>/dev/null || echo dev)

REVISION = "${TAG}:${COMMIT}-${ITERATION}@${ENV}"

init:
	go mod init sockstat_exporter

vet:
	go vet -mod=vendor $(shell go list ./...)

deps:
	go mod tidy
	go mod vendor

run:
	go run -mod=vendor . --metrics.go=false

build:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -o sockstat_exporter -ldflags "-s -w -X main.revision=${REVISION}"
