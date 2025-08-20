FROM golang:1.25rc3-alpine3.22

LABEL authors="marckent04"

WORKDIR /app

RUN apk add --no-cache curl

COPY . .

EXPOSE 3000

CMD go run main.go