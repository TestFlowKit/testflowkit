FROM golang:1.24.0-alpine

LABEL authors="marckent04"

WORKDIR /app

# for healthcheck
RUN apk add --no-cache curl

COPY . .

CMD go run main.go