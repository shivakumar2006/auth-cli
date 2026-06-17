FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-cli .

FROM alpine 

WORKDIR /app

COPY --from=builder /app/auth-cli .
COPY --from=builder /app/migrations ./migrations

RUN mkdir -p qrcodes

CMD ["./auth-cli"]