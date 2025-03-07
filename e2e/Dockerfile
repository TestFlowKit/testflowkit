FROM golang:1.23.3-alpine AS build

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

ENTRYPOINT /app/tkit run -l ./features -c ./cli.yml -f ./frontend.yml --timeout 10s