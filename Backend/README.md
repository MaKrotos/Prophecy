# Prophecy Backend

Backend сервис для приложения Prophecy, написанный на языке Go.

## Структура проекта

```
Backend/
├── auth/            # Аутентификация и проверка токенов
├── config/          # Конфигурация приложения
├── database/        # Подключение к базе данных
├── handlers/        # Обработчики HTTP запросов
├── migrations/      # Миграции базы данных
├── models/          # Модели данных
├── routes/          # Определение маршрутов API
├── main.go          # Точка входа в приложение
├── go.mod           # Модуль Go и зависимости
├── go.sum           # Контрольные суммы зависимостей
└── Dockerfile       # Конфигурация Docker контейнера
```

## Зависимости

- Go 1.21 или выше
- Gin Web Framework
- PostgreSQL driver (github.com/lib/pq)
- JWT library (github.com/golang-jwt/jwt/v5)

## Установка и запуск

### Локальный запуск

1. Убедитесь, что у вас установлен Go 1.21 или выше
2. Перейдите в директорию Backend:
   ```
   cd Backend
   ```
3. Загрузите зависимости:
   ```
   go mod tidy
   ```
4. Установите goose для миграций:
   ```
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```
5. Запустите миграции (убедитесь, что PostgreSQL запущен):
   ```
   goose -dir migrations postgres "user=prophecy_user password=prophecy_password dbname=prophecy_db sslmode=disable" up
   ```
6. Запустите приложение:
   ```
   go run main.go
   ```

### Запуск через Docker

Приложение может быть запущено через Docker используя docker-compose из корневой директории проекта:

```
docker-compose up backend
```

## API Endpoints

- `GET /` - Приветственное сообщение
- `GET /health` - Проверка состояния сервиса
- `GET /users/:id` - Получение информации о пользователе по ID
- `POST /users` - Создание нового пользователя
- `POST /auth/telegram` - Проверка токена Telegram WebApp
- `GET /auth/telegram/token` - Получение токена Telegram бота (для тестирования)
- `GET /auth/verify` - Проверка JWT токена (требует заголовок Authorization: Bearer <token>)

## Аутентификация Telegram WebApp

Для проверки токена Telegram WebApp используется алгоритм HMAC-SHA256.
Токен бота должен быть установлен в переменной окружения `TELEGRAM_BOT_TOKEN`.

### Формат запроса для проверки токена:

```json
{
  "id": 123456789,
  "first_name": "John",
  "last_name": "Doe",
  "username": "johndoe",
  "photo_url": "https://t.me/i/userpic/320/johndoe.jpg",
  "auth_date": 1612345678,
  "hash": "abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
}
```

### Ответ при успешной аутентификации:

```json
{
  "message": "Token is valid",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "telegram_id": 123456789,
    "first_name": "John",
    "last_name": "Doe",
    "username": "johndoe",
    "photo_url": "https://t.me/i/userpic/320/johndoe.jpg",
    "auth_date": 1612345678,
    "created_at": "2025-10-09T14:40:10Z"
  }
}
```

## JWT Аутентификация

После успешной аутентификации через Telegram, сервер возвращает JWT токен, который можно использовать для авторизации в других endpoint'ах.

### Использование JWT токена

Для доступа к защищенным endpoint'ам, добавьте заголовок Authorization:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Проверка JWT токена

Endpoint `GET /auth/verify` проверяет валидность JWT токена и возвращает информацию о пользователе.

## Переменные окружения

- `SERVER_PORT` - Порт для запуска сервера (по умолчанию: 8080)
- `TELEGRAM_BOT_TOKEN` - Токен Telegram бота для проверки аутентификации
- `DB_HOST` - Хост базы данных (по умолчанию: localhost)
- `DB_PORT` - Порт базы данных (по умолчанию: 5432)
- `DB_USER` - Пользователь базы данных (по умолчанию: user)
- `DB_PASSWORD` - Пароль базы данных (по умолчанию: password)
- `DB_NAME` - Имя базы данных (по умолчанию: prophecy)
- `JWT_SECRET` - Секретный ключ для подписи JWT токенов (по умолчанию: prophecy_jwt_secret_key)

