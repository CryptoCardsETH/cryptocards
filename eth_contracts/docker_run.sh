#!/usr/bin/env bash
case $1 in
    runandtest)
    truffle migrate --reset
	  truffle test
    node gatherAddresses.js
    ;;
  *)
    exec "$@"
    ;;
esac
