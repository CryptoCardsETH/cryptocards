#!/usr/bin/env bash
case $1 in
    run)
		./proto_tools/bin/protoc --php_out=/backend/proto \
	   --grpc_out=/backend/proto \
	   --plugin=protoc-gen-grpc=grpc/bins/opt/grpc_php_plugin \
	   --proto_path=/ethereum_proxy/src/GoRpc/rpcServer/ \
	   /ethereum_proxy/src/GoRpc/rpcServer/rpcServer.proto;
        ;;
    *)
        exec "$@"
        ;;
esac
