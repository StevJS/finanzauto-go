FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/api/main.go"]