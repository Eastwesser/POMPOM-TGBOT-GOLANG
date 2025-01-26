# Используем базовый образ
FROM golang:1.21 as builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Сборка проекта
RUN go build -o pompon-bot ./cmd/bot/main.go

# Минимальный образ для запуска
FROM gcr.io/distroless/base-debian11

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированный файл
COPY --from=builder /app/pompon-bot .

# Открываем порт для бота (опционально)
EXPOSE 8080

# Запускаем приложение
CMD ["./pompon-bot"]
