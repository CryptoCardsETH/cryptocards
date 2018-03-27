#!/usr/bin/env bash

echo "THIS SCRIPT DEPRECATED IN FAVOR OF DOCKER"

cd "$(dirname "$0")"
echo "generating PHP via protoc..."
#php:
protoc --php_out=../../backend/proto \
	   --grpc_out=../../backend/proto \
	   --plugin=protoc-gen-grpc=grpc/bins/opt/grpc_php_plugin \
	   --proto_path=../../ethereum_proxy/src/GoRpc/rpcServer/ \
	   ../../ethereum_proxy/src/GoRpc/rpcServer/rpcServer.proto;


echo "generating golang via protoc..."

cd "$PWD/../../ethereum_proxy/src/GoRpc";
protoc -I rpcServer/ rpcServer/rpcServer.proto \
	   --go_out=plugins=grpc:rpcServer;
