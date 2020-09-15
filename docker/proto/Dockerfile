FROM golang:1.14.0

ENV DEBIAN_FRONTEND=noninteractive

ARG PROTO_VERSION=3.12.4

WORKDIR /proto

COPY ./proto .

RUN mkdir /output /output/server /output/client

RUN apt-get -qq update && apt-get -qq install -y \
  unzip

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/protoc-${PROTO_VERSION}-linux-x86_64.zip -o protoc.zip && \
  unzip -qq protoc.zip && \
  cp ./bin/protoc /usr/local/bin/protoc && \
  cp -r ./include /usr/local

RUN go get -u github.com/golang/protobuf/protoc-gen-go