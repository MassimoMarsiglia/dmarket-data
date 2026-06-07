FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /dmarket-bot main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /dmarket-bot /app/dmarket-bot

CMD ["/app/dmarket-bot", "run", "--cfg", "/app/bin/cfg.yaml"]
