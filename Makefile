.PHONY: all

prepare:
	brew install protoc

	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/protoc-gen-go
