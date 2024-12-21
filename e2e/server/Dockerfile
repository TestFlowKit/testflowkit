FROM golang:1.23.3-alpine

LABEL authors="marckent04"

WORKDIR /app

# for healthcheck
RUN apk add --no-cache curl

COPY . .

CMD go run main.go