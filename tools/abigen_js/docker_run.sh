#!/usr/bin/env bash
case $1 in
    run)
        node /tools/abigen_js/abigen.js 
	  ;;
    *)
        exec "$@"
        ;;
esac
