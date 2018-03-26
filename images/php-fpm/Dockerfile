FROM php:7.1-fpm

RUN apt-get update && \
    docker-php-ext-install pdo_mysql && \
    apt-get install -y curl git zlib1g-dev

RUN pecl install grpc\
    && docker-php-ext-enable grpc 

RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer
RUN composer self-update