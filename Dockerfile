FROM golang:1.18-alpine as builder

ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /go/src/ip-addr

COPY . /go/src/ip-addr/
RUN CGO_ENABLED=0 go install -ldflags "-s -w"

FROM scratch

# install ip-addr
COPY --from=builder /go/bin/ip-addr /usr/local/bin/ip-addr

ENTRYPOINT [ "ip-addr" ]

EXPOSE 80
