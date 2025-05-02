FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go mod download
RUN go mod tidy

CMD ["sh", "-c", "go run ./cmd/app/main.go"]