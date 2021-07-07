FROM golang:1.16 as builder
WORKDIR /go/go-echo
COPY . .

ARG BUILD_VERSION
ENV BUILD_VERSION=${BUILD_VERSION:-v0.1.0}
RUN make clean && make build

FROM alpine:3.14
WORKDIR /opt/go-echo
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/go-echo/bin/echo echo
ENTRYPOINT ["./echo", "serve"]
