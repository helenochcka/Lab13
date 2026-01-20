FROM golang:1.25.2

WORKDIR /app

COPY event-service/go.mod event-service/go.sum ./

RUN go mod download

COPY event-service ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /event-service ./cmd/main.go

CMD ["/event-service"]
