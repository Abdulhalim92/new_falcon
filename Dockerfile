# Build stage
FROM golang:1.21-alpine AS builder

ENV GOOS=linux

WORKDIR /app

COPY . .
COPY go.mod go.sum ./

RUN go mod download

RUN go build -o falconApi main.go

# Run stage
FROM alpine

WORKDIR /app

COPY --from=builder /app/falconApi .
COPY --from=builder /app/config .

EXPOSE 8003

CMD ["/app/falconApi"]