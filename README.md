# Prophecy - Проект

## Описание

Prophecy - это веб-приложение с фронтендом на Vue.js, бэкендом на Go и базой данных PostgreSQL. Проект использует Docker и Docker Compose для удобного развертывания.

## Требования

- Docker
- Docker Compose

## Быстрый запуск (Production)

Для запуска проекта в production режиме выполните:

```bash
docker-compose up -d
```

Эта команда запустит все необходимые сервисы:

- Фронтенд приложение на порту 80 (HTTP)
- Бэкенд API на порту 8080
- Базу данных PostgreSQL на порту 5432

## Пример упрощенного docker-compose для быстрого запуска

```yaml
services:

  traefik:
    image: traefik:v2.10
    container_name: traefik
    restart: unless-stopped
    command:
      - --api.insecure=true
      - --providers.docker=true
      - --providers.docker.exposedbydefault=false
      - --entrypoints.web.address=:80
      - --entrypoints.websecure.address=:443
      - --certificatesresolvers.myresolver.acme.tlschallenge=true
      - --certificatesresolvers.myresolver.acme.email=your-email@example.com
      - --certificatesresolvers.myresolver.acme.storage=/letsencrypt/acme.json
    ports:
      - "80:80"
      - "443:443"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./letsencrypt:/letsencrypt
    networks:
      - Prophecy-network

  frontend:
    build:
      context: .
      dockerfile: Dockerfile
    # image: ip/makrotos/prophecy:latest
    restart: unless-stopped
    pull_policy: always
    container_name: frontend
    environment:
      - NODE_ENV=production
    networks:
      - Prophecy-network
    depends_on:
      - backend
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.frontend.rule=Host(`your-domain.com`)"
      - "traefik.http.routers.frontend.entrypoints=websecure"
      - "traefik.http.routers.frontend.tls.certresolver=myresolver"
      - "traefik.http.services.frontend.loadbalancer.server.port=80"

  watchtower:
    image: containrrr/watchtower
    container_name: watchtower
    restart: unless-stopped
    environment:
      - WATCHTOWER_CLEANUP=true
      - WATCHTOWER_POLL_INTERVAL=300
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - Prophecy-network

  backend:
    build:
      context: ./Backend
      dockerfile: Dockerfile
    container_name: backend-prophecy
    restart: unless-stopped
    pull_policy: always
    environment:
      - SERVER_PORT=8080
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_INTERNAL_PORT=5432
      - DB_USER=example_user
      - DB_PASSWORD=example_password
      - DB_NAME=example_db
      - JWT_SECRET=example_jwt_secret
      - TELEGRAM_BOT_TOKEN=1234567890:ABC_DEF_GHI_JKL_MNO_PQR_STU
    # - ADMIN_TELEGRAM_ID=123456789
    networks:
      - Prophecy-network
    depends_on:
      - postgres

  postgres:
    image: postgres:16
    container_name: postgres-prophecy
    restart: unless-stopped
    environment:
      - POSTGRES_DB=example_db
      - POSTGRES_USER=example_user
      - POSTGRES_PASSWORD=example_password
    ports:
      - "5432:5432"
    networks:
      - Prophecy-network
    volumes:
      - pgdata:/var/lib/postgresql/data

networks:
  Prophecy-network:
    driver: bridge

volumes:
  pgdata:
```

## Разработка

Для запуска проекта в режиме разработки выполните:

```bash
docker-compose up frontend-dev-Prophecy
```

Фронтенд для разработки будет доступен по адресу: http://localhost:5173

Для запуска всех сервисов в режиме разработки:

```bash
docker-compose up
```

## Сервисы

| Сервис       | Порт(ы) | Описание                     |
| ------------ | ------- | ---------------------------- |
| frontend     | 80      | Основное фронтенд приложение |
| backend      | 8080    | API бэкенд                   |
| postgres     | 5432    | База данных PostgreSQL       |
| frontend-dev | 5173    | Фронтенд для разработки      |

## Переменные окружения

### Frontend

| Переменная | Значение по умолчанию | Описание                |
| ---------- | --------------------- | ----------------------- |
| NODE_ENV   | production            | Режим работы приложения |

### Frontend Development

| Переменная          | Значение по умолчанию | Описание                                           |
| ------------------- | --------------------- | -------------------------------------------------- |
| NODE_ENV            | development           | Режим работы приложения                            |
| CHOKIDAR_USEPOLLING | true                  | Использовать опрос файловой системы для hot-reload |

### Backend

| Переменная         | Значение по умолчанию                        | Описание                        |
| ------------------ | -------------------------------------------- | ------------------------------- |
| SERVER_PORT        | 8080                                         | Порт сервера                    |
| DB_HOST            | postgres                                     | Хост базы данных                |
| DB_PORT            | 5432                                         | Порт базы данных                |
| DB_INTERNAL_PORT   | 5432                                         | Внутренний порт базы данных     |
| DB_USER            | prophecy_user                                | Пользователь базы данных        |
| DB_PASSWORD        | prophecy_password                            | Пароль пользователя базы данных |
| DB_NAME            | prophecy_db                                  | Название базы данных            |
| JWT_SECRET         | prophecy_jwt_secret_key_change_in_production | Секретный ключ для JWT          |
| TELEGRAM_BOT_TOKEN | 123456789:ABCDEFabcdef1234567890ABCDEFabcd   | Токен Telegram бота             |
| ADMIN_TELEGRAM_ID  | 123456789                                    | ID администратора в Telegram    |

**Важно:** Для production использования обязательно измените следующие переменные:

- `JWT_SECRET` - установите уникальный секретный ключ
- `TELEGRAM_BOT_TOKEN` - установите токен вашего Telegram бота
- `ADMIN_TELEGRAM_ID` - установите ID вашего Telegram аккаунта

### PostgreSQL

| Переменная        | Значение по умолчанию | Описание                        |
| ----------------- | --------------------- | ------------------------------- |
| POSTGRES_DB       | prophecy_db           | Название базы данных            |
| POSTGRES_USER     | prophecy_user         | Пользователь базы данных        |
| POSTGRES_PASSWORD | prophecy_password     | Пароль пользователя базы данных |

## Остановка проекта

Для остановки всех сервисов выполните:

```bash
docker-compose down
```

Для остановки с удалением томов (включая данные базы):

```bash
docker-compose down -v
```
