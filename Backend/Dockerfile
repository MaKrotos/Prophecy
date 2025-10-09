# Используем официальный образ Golang для сборки
FROM golang:1.23-alpine AS builder

# Установка переменных окружения
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Установка рабочей директории
WORKDIR /app

# Копирование файлов модулей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка бинарного файла
RUN go build -a -installsuffix cgo -o main .

# Установка goose для миграций
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Используем минимальный образ Alpine для запуска
FROM alpine:latest

# Установка рабочей директории
WORKDIR /root/

# Установка зависимостей для goose и SSL
RUN apk add --no-cache ca-certificates

# Создание директории для сертификатов
RUN mkdir -p /etc/letsencrypt/live/prophecy.makrotos.ru

# Копирование бинарного файла из builder образа
COPY --from=builder /app/main .
COPY --from=builder /go/bin/goose /usr/local/bin/goose

# Копирование миграций
COPY --from=builder /app/migrations ./migrations

# Открытие порта 8080
EXPOSE 8080

# Команда запуска приложения с миграциями
CMD ["sh", "-c", "goose -dir ./migrations postgres \"host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable\" up && ./main"]