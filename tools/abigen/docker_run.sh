#!/usr/bin/env bash
case $1 in
    run)
        abigen --sol=/eth_contracts/contracts/CryptoCardsCore.sol --pkg=contracts --out=/ethereum_proxy/src/GoRpc/contracts/core.go
        echo "Generate core.go"
	      ;;
    *)
        exec "$@"
        ;;
esac
