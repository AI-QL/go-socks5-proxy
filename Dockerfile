ARG GOLANG_VERSION="1.23.1"

FROM golang:$GOLANG_VERSION-alpine AS builder
RUN apk --no-cache add tzdata
WORKDIR /go/src/go-socks5-proxy/socks5
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./socks5

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /go/src/go-socks5-proxy/socks5/socks5 /
ENTRYPOINT ["/socks5"]