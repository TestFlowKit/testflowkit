FROM golang:1.25.0-bookworm AS doc_generation

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN make generate_doc

FROM node:22-bookworm AS builder

WORKDIR /app

COPY --from=doc_generation /app/documentation .

RUN yarn

RUN yarn build


FROM gcr.io/distroless/nodejs22-debian12

WORKDIR /app

COPY --from=builder /app/.output .

EXPOSE 3000

CMD ["server/index.mjs"]

