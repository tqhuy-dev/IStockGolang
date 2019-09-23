FROM golang:1.13-alpine as builder
WORKDIR /app
COPY go.mod  .
COPY go.sum .

# Download dependencies

RUN go mod Download
COPY . .

FROM alpine:3.5
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
 cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
 apk del tzdata

WORKDIR /app

RUN go build -o main
ENTRYPOINT ["./main"]
