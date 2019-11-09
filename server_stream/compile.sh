#!/bin/sh

echo "compile go stub file"
protoc --proto_path=./proto --go_out=plugins=grpc:./proto ./proto/greeting-service.proto

# echo "compile java stub file"
# protoc --proto_path=./proto --java_out=plugins=grpc:./proto ./proto/todo-service.proto
