#!/bin/bash

for i in `find api -name *.proto|xargs`;do
    protoc --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
done

# 服务器
go run lithum.go -conf configs.yaml

# 客户端
go run lithum.go -conf grpc://127.0.0.1:9090?uuid=061aabc4-ed98-4d4c-8b3e-413bf181e3e1