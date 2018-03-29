#!/usr/bin/env bash
case $1 in
    run)
         go run GoRpc/server/main.go
        ;;
    *)
        exec "$@"
        ;;
esac
