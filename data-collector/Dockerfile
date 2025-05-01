FROM golang:1.21

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y postgresql postgresql-contrib && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY . /app

RUN go mod download
RUN go mod tidy

EXPOSE 8080

CMD ["sh", "-c", "go run ./cmd/app/main.go"]