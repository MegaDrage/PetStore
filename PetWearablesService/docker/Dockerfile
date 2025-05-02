FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o pet-wearables ./cmd/pet-wearables

FROM alpine:3.21.3
WORKDIR /root/
COPY --from=builder /app/pet-wearables .
COPY .env .
EXPOSE 8085
CMD ["./pet-wearables"]