FROM golang:1.25.1-bookworm AS build

WORKDIR /app

COPY . .

WORKDIR /app/cmd/testflowkit

RUN go build -o tkit


# FROM debian:bookworm-slim
FROM golang:1.25.1-bookworm


ARG DRIVER
ARG INSTALL_CMD
ARG TAGS

LABEL authors="marckent04"

WORKDIR /app

COPY ./e2e .
COPY --from=build /app/cmd/testflowkit/tkit .


RUN echo "driver: \"$DRIVER\"" >> ./env.ci.yml && \
	eval "$INSTALL_CMD"

ENV tags=$TAGS

ENTRYPOINT ["/bin/sh", "-c", "/app/tkit run --tags=${tags} --env-file=./env.ci.yml"]

