#!/usr/bin/env bash

echo "THIS SCRIPT DEPRECATED IN FAVOR OF DOCKER"

cd "$(dirname "$0")"
git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc
cd grpc
git submodule update --init
make grpc_php_plugin