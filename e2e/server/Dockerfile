FROM golang:1.25.0-alpine

LABEL authors="marckent04"

WORKDIR /app

RUN apk add --no-cache curl

COPY . .

EXPOSE 3000

CMD go run main.go