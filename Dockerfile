# syntax=docker/dockerfile:1.7
FROM golang:1.26-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git bash ca-certificates openssh

ENV GONOSUMDB=github.com/31c0cd03/petition-pkg

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    if [ -f "go.work" ]; then \
    go mod download -C ./internal/svc/gateway; \
    else \
    go mod download; \
    fi

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    if [ -f "go.work" ]; then \
    go build -o /app/gateway-svc ./internal/svc/gateway/cmd/main.go; \
    else \
    go build -o /app/gateway-svc ./cmd/main.go; \
    fi

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/gateway-svc ./gateway-svc

EXPOSE 8000
CMD ["./gateway-svc"]