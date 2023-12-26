FROM golang:1.21

WORKDIR /app

COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -o golang-crud .

EXPOSE 8080

CMD ["./golang-crud"]
