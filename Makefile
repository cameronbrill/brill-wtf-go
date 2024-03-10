PROTO_DIR=grpc
BUILD_DIR=build
SUB_DIRS=example grpc/server rest
BUILD_TARGETS=$(addprefix cmd/,$(SUB_DIRS))
SHELL := /bin/bash

all: deps $(BUILD_TARGETS)
$(BUILD_TARGETS): %: 
	go build -o '$(BUILD_DIR)/$(subst cmd-,$e,$(subst /,-,$@))' '$@'/main.go
.PHONY: deps $(BUILD_TARGETS)

rest: deps cmd/rest
	go build -o tmp/rest ./cmd/rest/main.go

deps: 
	go mod tidy
	go get

lint:
	golangci-lint run --config .golangci.yml

grpc/generate:
	protoc -Igrpc \
		--go_out=./${PROTO_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=./${PROTO_DIR} --go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/*.proto

docker-build: 
	DOCKER_BUILDKIT=1 docker build -t brill.wtf .

docker-run:
	docker run \
		-it --rm \
		-p 3333:3333 \
		--env-file <(doppler secrets download --no-file --format docker) \
		brill.wtf

docker: docker-build docker-run

clean:
	rm build/*
