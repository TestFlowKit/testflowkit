FROM golang:1.24.0-alpine AS build

WORKDIR /app

COPY . .

WORKDIR /app/cmd/testflowkit

RUN go build -o tkit


FROM alpine:3.21

LABEL authors="marckent04"

WORKDIR /app

COPY ./e2e .

COPY --from=build /app/cmd/testflowkit .

RUN apk add --no-cache chromium

ENTRYPOINT /app/tkit run --location="./features" --timeout="10s" --env="ci"