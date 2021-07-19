#!/bin/bash

#for i in `find api -name *.proto|xargs`;do
#    protoc --proto_path=. --proto_path=third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. $i
#    protoc --proto_path=. --proto_path=third_party --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. $i
#done

#cd app/file

#for i in `find . -name *.proto|xargs`;do
#    protoc --proto_path=. --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. $i
#done


#cd internal/data
#ent generate ./ent/schema
#cd ../..

#cd cmd/server
#wire
#cd ../..

#mkdir -p bin
#go build -ldflags "-X main.Version=0.0.1" -o ./bin/ ./...
#./bin/server -conf configs/config.yaml

go run lithum.go