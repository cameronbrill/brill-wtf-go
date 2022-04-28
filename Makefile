BINARY_NAME=go-project-template
PROTO_DIR=grpc

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=window go build -o ${BINARY_NAME}-windows main.go

generate:
	protoc -Igrpc \
		--go_out=./${PROTO_DIR} --go_opt=paths=source_relative \
		--go-grpc_out=./${PROTO_DIR} --go-grpc_opt=paths=source_relative \
		${PROTO_DIR}/*.proto

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows