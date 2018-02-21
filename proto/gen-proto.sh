#!/usr/bin/env bash
cd "$(dirname "$0")"
echo "generating PHP via protoc..."
# php:
protoc --php_out=../backend/proto \
	   --grpc_out=../backend/proto \
	   --plugin=protoc-gen-grpc=grpc/bins/opt/grpc_php_plugin \
	   --proto_path=../go/src/GoRpc/rpcServer/ \
	   ../go/src/GoRpc/rpcServer/rpcServer.proto;

# protoc --php_out=out --grpc-php_out=composer_name=grpc/hello:out --plugin=protoc-gen-grpc-php=./vendor/bin/protoc-gen-grpc-php ./helloworld.proto

echo "generating golang via protoc..."

cd "$PWD/../go/src/GoRpc";
protoc -I rpcServer/ rpcServer/rpcServer.proto \
	   --go_out=plugins=grpc:rpcServer;