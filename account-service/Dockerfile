FROM golang:1.25.2

WORKDIR /app

COPY account-service/go.mod account-service/go.sum ./

RUN go mod download

COPY account-service ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /account-service ./cmd/main.go

CMD ["/account-service"]
