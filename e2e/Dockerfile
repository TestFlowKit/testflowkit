FROM golang:1.23.3-alpine AS build

WORKDIR /app

COPY . .

WORKDIR /app/cmd/etoolse

RUN go build -o etoolse


FROM alpine:3.21

LABEL authors="marckent04"

WORKDIR /app

COPY ./e2e .

COPY --from=build /app/cmd/etoolse .

RUN apk add --no-cache chromium

ENTRYPOINT /app/etoolse run -l ./features -c ./cli.yml -f ./frontend.yml --timeout 10s