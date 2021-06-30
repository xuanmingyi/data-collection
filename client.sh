#!/bin/bash

cd client
ent generate ./ent/schema
mkdir -p bin/ &&go build -o bin/ ./...
./bin/client