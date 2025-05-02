FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY cmd/simulate_collar/main.go cmd/simulate_collar/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o /collar-simulator cmd/simulate_collar/main.go

FROM alpine:3.21.3

COPY --from=builder /collar-simulator /collar-simulator

ENTRYPOINT ["/collar-simulator"]