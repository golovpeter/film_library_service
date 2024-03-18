FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/service

EXPOSE 8080

CMD ["./main"]