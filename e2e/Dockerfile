FROM golang:1.25.0-alpine AS build

WORKDIR /app

COPY . .

WORKDIR /app/cmd/testflowkit

RUN go build -o tkit


# FROM node:24-alpine3.22
FROM ubuntu:24.04

LABEL authors="marckent04"

WORKDIR /app

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    gnupg \
    && rm -rf /var/lib/apt/lists/*


RUN curl -fsSL https://deb.nodesource.com/setup_24.x | bash - \
    && apt-get install -y nodejs \
    && rm -rf /var/lib/apt/lists/*

RUN npx playwright install --with-deps

COPY ./e2e .

COPY --from=build /app/cmd/testflowkit .

# RUN apk add --no-cache chromium

ENTRYPOINT /app/tkit run --location="./features" --timeout="10s" --env-file="./env.ci.yml"
