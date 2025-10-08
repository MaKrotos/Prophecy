#!/bin/sh

# Генерируем финальный nginx.conf из шаблона
export DOMAIN=${DOMAIN:-localhost}
envsubst < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Запускаем nginx
nginx

# Запускаем Vite dev-сервер
npm run dev -- --host 0.0.0.0 --port 80 