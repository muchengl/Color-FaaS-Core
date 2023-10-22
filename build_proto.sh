#!/usr/bin/env bash
export PATH="$PATH:$(go env GOPATH)/bin"

echo "  ---> make clean <---"
rm -rf pkg/proto/executor/*.go
rm -rf pkg/proto/scheduler/*.go

echo "  ---> make executor.proto <---"
protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
      pkg/proto/executor/executor.proto

echo "  ---> make scheduler.proto <---"
protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
      pkg/proto/scheduler/scheduler.proto