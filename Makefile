PROTO_DIR=grpc
BUILD_DIR=build
SUB_DIRS=example grpc/server rest
BUILD_TARGETS=$(addprefix cmd/,$(SUB_DIRS))

all: deps $(BUILD_TARGETS)
$(BUILD_TARGETS): %: 
	go build -o '$(BUILD_DIR)/$(subst cmd-,$e,$(subst /,-,$@))' '$@'/main.go
.PHONY: deps $(BUILD_TARGETS)

rest: deps cmd/rest
	go build -o tmp/rest ./cmd/rest/main.go

deps: 
	go mod tidy
	go get

grpc/generate:
	protoc -Igrpc \
		--go_out=./${PROTO_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=./${PROTO_DIR} --go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/*.proto

clean:
	rm build/*
