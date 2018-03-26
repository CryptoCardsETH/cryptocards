#!/usr/bin/env bash
case $1 in
  runandtest)
    truffle migrate
	truffle test
    ;;
  *)
    exec "$@"
    ;;
esac
