FROM golang:1.20-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    git \
    && update-ca-certificates

# build stage
FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod download && go mod verify

RUN go build -o users-api -a .

# prod stage
FROM alpine:latest as prod

COPY --from=builder /app/users-api /usr/local/bin/users-api
EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/users-api"]
