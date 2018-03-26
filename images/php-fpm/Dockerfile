FROM php:fpm

RUN apt-get update && docker-php-ext-install pdo_mysql && apt-get install -y curl git

RUN curl -sS https://getcomposer.org/installer | php
RUN mv composer.phar /usr/local/bin/composer
RUN composer self-update