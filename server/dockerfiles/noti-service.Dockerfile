#build stage
FROM golang:1.20 AS builder
WORKDIR /go/src/capsmhoo
ADD . .
RUN go get capsmhoo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app cmd/noti/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/capsmhoo/app .
COPY --from=builder /go/src/capsmhoo/config/config-compose.yaml ./config/config.yaml
CMD ["./app"]

EXPOSE 50054