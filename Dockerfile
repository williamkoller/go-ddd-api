FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /api ./cmd/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /api .

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 3003

CMD ["./api"]
