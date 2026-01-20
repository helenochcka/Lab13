FROM golang:1.25.2

WORKDIR /app

COPY notification-service/go.mod notification-service/go.sum ./

RUN go mod download

COPY notification-service ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /notification-service ./cmd/main.go

CMD ["/notification-service"]
