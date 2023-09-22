#build stage
FROM golang:1.20 AS builder
WORKDIR /go/src/capsmhoo
ADD . .
RUN go get capsmhoo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app cmd/team/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/capsmhoo/app .
COPY --from=builder /go/src/capsmhoo/config/config.yaml ./config/
CMD ["./app"]

EXPOSE 50051