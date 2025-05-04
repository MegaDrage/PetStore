FROM php:8.3-fpm

RUN apt-get update && apt-get install -y \
    git \
    unzip \
    libpq-dev \
    libzip-dev \
    zlib1g-dev \
    libicu-dev \
    g++ \
    libssl-dev \
    pkg-config \
    && docker-php-ext-configure intl \
    && docker-php-ext-install pdo zip intl \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN pecl install mongodb && docker-php-ext-enable mongodb

COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

WORKDIR /var/www/html

COPY . /var/www/html

RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html

RUN mkdir -p var/cache && chmod -R 777 var/cache

RUN composer install --no-interaction --optimize-autoloader

COPY docker/php/php.ini /usr/local/etc/php/php.ini

CMD ["php", "-S", "0.0.0.0:8000", "-t", "public"]