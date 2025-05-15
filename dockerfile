From golang:1.23

WORKDIR /usr/src/app

# Устанавливаем git, чтобы Go мог скачать модули
RUN apt-get update && apt-get install -y git

# Устанавливаем Air, теперь по новому модульному пути
RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy
