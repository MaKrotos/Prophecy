#!/bin/sh

SSL_CERT="/etc/letsencrypt/live/fullchain.pem"
SSL_KEY="/etc/letsencrypt/live/privkey.pem"

# Удаляем все .conf кроме нужного
rm -f /etc/nginx/conf.d/*.conf

if [ -f "$SSL_CERT" ] && [ -f "$SSL_KEY" ]; then
  echo "SSL сертификаты найдены, включаю HTTPS"
  cp /nginx-tpl/https.conf /etc/nginx/conf.d/default.conf
else
  echo "SSL сертификаты не найдены, запускаю только HTTP"
  cp /nginx-tpl/http.conf /etc/nginx/conf.d/default.conf
fi

nginx -g 'daemon off;'