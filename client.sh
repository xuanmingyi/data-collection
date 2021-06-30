#!/bin/bash

for i in `find api -name *.proto|xargs`;do
    protoc --proto_path=. --proto_path=third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
done

cd client

for i in `find . -name *.proto|xargs`;do
    protoc --proto_path=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
done

ent generate ./ent/schema
mkdir -p bin/ &&go build -o bin/ ./...
./bin/client