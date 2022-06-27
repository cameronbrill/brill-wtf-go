# syntax=docker/dockerfile:1
FROM golang:1.18-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
ADD internal/* /app
RUN go mod download

ADD . /app

RUN go build -o tmp/rest cmd/rest/main.go

FROM akitasoftware/cli:latest

WORKDIR /root/

COPY --from=build /app/tmp/rest ./rest

EXPOSE 3333

CMD [ "apidump", "--service", "brill-wtf", "--filter", "port 3333", "-u", "root", "-c", "./rest" ]