FROM golang:alpine AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn go build -ldflags "-X main.Version=1.0.0" -o main main.go

FROM alpine:latest

COPY --from=builder /src/main /app/main

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

ENTRYPOINT ["./main"]
