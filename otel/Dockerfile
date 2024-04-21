FROM golang:1.22 AS builder


WORKDIR /src
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn go mod download
COPY . .

RUN make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["./otel", "-conf", "/data/conf"]
