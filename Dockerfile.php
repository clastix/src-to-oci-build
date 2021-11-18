ARG VERSION=8.0

FROM php:${VERSION}-apache-bullseye

COPY . /var/www/html

EXPOSE 80
