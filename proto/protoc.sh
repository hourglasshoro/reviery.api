#!/bin/sh

#set -xe
#
#SERVER_OUTPUT_DIR=/output/server
#CLIENT_OUTPUT_DIR=/output/client
#
#protoc --version
#protoc -I=/proto/protos customer.proto staff.proto kitchen.proto\
#  --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}" \
#  --dart_out="grpc:${CLIENT_OUTPUT_DIR}"
#
#protoc -I=/proto/protos timestamp.proto wrappers.proto\
#  --dart_out="grpc:${CLIENT_OUTPUT_DIR}"