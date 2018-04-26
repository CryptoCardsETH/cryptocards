#!/usr/bin/env bash
case $1 in
    run)
         #go get -v ./...
        go get -v github.com/fschl/CompileDaemon
        cd $GOPATH/src/github.com/fschl/CompileDaemon \
            && go install
        #go run GoRpc/server/main.go GoRpc/server/helpers.go
        cd /ethereum_proxy/src
        ../bin/CompileDaemon -build="go build GoRpc/server" -recursive=true -command="go run GoRpc/server/main.go GoRpc/server/helpers.go" -color=true
        ;;
    *)
        exec "$@"
        ;;
esac
