# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /itx-git-server

##
## Deploy
##
FROM ubuntu
RUN apt-get -y update
RUN apt-get -y install git

WORKDIR /

COPY --from=build /itx-git-server /itx-git-server
COPY config.yaml /config.yaml

VOLUME /opt

EXPOSE 5000
ENTRYPOINT ["/itx-git-server"]