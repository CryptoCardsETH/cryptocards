#!/bin/bash

case "$1" in
    backend-migrate)
        docker-compose exec fpm php artisan migrate
        ;;
    backend-tail)
        docker-compose exec fpm php artisan tail
        ;;
    backend-init)
        docker-compose exec fpm php artisan jwt:secret
        docker-compose exec fpm php artisan key:generate
        docker-compose exec fpm php artisan migrate
        ;;
    *)
        echo "No command specified!"
        exit 1
esac
