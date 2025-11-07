# Stage 1: Build (production build)
FROM node:lts-alpine AS build
WORKDIR /app

# Копируем package.json и package-lock.json
COPY package*.json ./

# Устанавливаем все зависимости (включая devDependencies для сборки)
RUN npm ci

# Копируем всё остальное (исходный код)
COPY . .

# Собираем приложение
RUN npm run build

# Удаляем devDependencies для уменьшения размера образа
RUN npm prune --production

# Stage 2: Production (nginx)
FROM nginx:alpine AS prod
# Копируем собранный фронт
COPY --from=build /app/dist /usr/share/nginx/html
# Копируем упрощенный nginx.conf
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]