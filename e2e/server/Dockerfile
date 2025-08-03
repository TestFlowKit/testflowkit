FROM golang:1.24.0-alpine

LABEL authors="marckent04"

WORKDIR /app

# for healthcheck
RUN apk add --no-cache curl

COPY . .

# Make healthcheck script executable
RUN chmod +x healthcheck.sh

CMD go run main.go