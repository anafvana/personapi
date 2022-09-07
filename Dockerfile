FROM golang:latest

EXPOSE 8080

WORKDIR /personapi
COPY . .
RUN go mod download

CMD go run startup/personapi/main.go