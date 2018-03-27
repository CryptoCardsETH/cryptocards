#!/usr/bin/env bash
case $1 in
    run)
		echo "running php codegen..."
		./proto_tools/bin/protoc --php_out=/backend/proto \
	   --grpc_out=/backend/proto \
	   --plugin=protoc-gen-grpc=grpc/bins/opt/grpc_php_plugin \
	   --proto_path=/ethereum_proxy/src/GoRpc/rpcServer/ \
	   /ethereum_proxy/src/GoRpc/rpcServer/rpcServer.proto;
	  echo "backend code updated from rpcServer.proto (/backend/proto)"

     echo "running golang codegen"
     cd /ethereum_proxy/src/GoRpc;
		 /tmp/proto_tools/bin/protoc -I rpcServer/ rpcServer/rpcServer.proto \
	   --go_out=plugins=grpc:rpcServer;
     ;;
    *)
        exec "$@"
        ;;
esac
