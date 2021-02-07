ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine AS builder
RUN apk add --update --no-cache ca-certificates make git curl gcc libc-dev
RUN mkdir -p /build
WORKDIR /build
COPY . /build/
RUN go mod download
RUN  make build-linux

FROM scratch
COPY --from=builder /build/rproxy /rproxy
ENTRYPOINT ["/rproxy"]
