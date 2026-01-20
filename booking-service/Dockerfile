FROM golang:1.25.2

WORKDIR /app

COPY booking-service/go.mod booking-service/go.sum ./

RUN go mod download

COPY booking-service ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /booking-service ./cmd/main.go

CMD ["/booking-service"]
