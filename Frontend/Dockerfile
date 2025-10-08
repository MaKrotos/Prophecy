# Stage 1: Build (production build)
FROM node:lts-alpine AS build
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем ВСЕ зависимости (включая devDependencies для сборки)
RUN npm ci --include=dev

# Копируем всё остальное (исходный код)
COPY . .
RUN npm run build

# Stage 2: Production (nginx)
FROM nginx:alpine AS prod
# Устанавливаем gettext для envsubst
RUN apk add --no-cache gettext
# Копируем собранный фронт
COPY --from=build /app/dist /usr/share/nginx/html
# Копируем шаблоны в отдельную папку
COPY nginx-tpl/http.conf /nginx-tpl/http.conf
COPY nginx-tpl/https.conf /nginx-tpl/https.conf
# Копируем кастомный entrypoint
COPY nginx-entrypoint.sh /nginx-entrypoint.sh
RUN chmod +x /nginx-entrypoint.sh
# Используем кастомный entrypoint
ENTRYPOINT ["/nginx-entrypoint.sh"]
EXPOSE 80
EXPOSE 443