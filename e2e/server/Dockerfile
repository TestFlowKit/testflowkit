FROM golang:1.25rc3-alpine3.22

LABEL authors="marckent04"

WORKDIR /app

# for healthcheck
RUN apk add --no-cache curl

COPY . .

# Make healthcheck script executable
RUN chmod +x healthcheck.sh

CMD go run main.go