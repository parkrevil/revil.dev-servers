#!/bin/sh

protoc libs/**/**/**.proto \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  --experimental_allow_proto3_optional


go get -d -u github.com/golang/protobuf/proto
go get -d -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go