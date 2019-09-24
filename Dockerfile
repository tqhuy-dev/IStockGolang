FROM golang:1.13-alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main
ENTRYPOINT ["./main"]
