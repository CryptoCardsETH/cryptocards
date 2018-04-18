#!/usr/bin/env bash
case $1 in
    run)
         #go get -v ./...
         go run GoRpc/server/main.go GoRpc/server/helpers.go
        ;;
    *)
        exec "$@"
        ;;
esac
