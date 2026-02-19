FROM golang:1.25.1-alpine AS build

WORKDIR /app

COPY . .

WORKDIR /app/cmd/testflowkit

RUN go build -o tkit


# FROM debian:bookworm-slim
FROM golang:1.25.1-alpine


ARG DRIVER
ARG INSTALL_CMD
ARG TAGS

LABEL authors="marckent04"

WORKDIR /app

COPY ./e2e .
COPY --from=build /app/cmd/testflowkit/tkit .


RUN echo "driver: \"$DRIVER\"" >> ./env.ci.yml

RUN $INSTALL_CMD
# RUN ./tkit install

ENV tags=$TAGS

ENTRYPOINT /app/tkit run --tags=$tags --env-file=./env.ci.yml

