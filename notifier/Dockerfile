FROM golang:1.21

WORKDIR /app

COPY . /app

RUN go mod download
RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "./cmd/app/main.go"]