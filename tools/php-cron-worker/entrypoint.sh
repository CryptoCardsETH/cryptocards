#!/bin/bash
# Run scheduler
while [ true ]
do
  php /var/www/laravel/artisan schedule:run
  sleep 60
done