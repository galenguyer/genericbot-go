FROM docker.io/library/golang:1.16.6-alpine3.14 as builder
LABEL org.opencontainers.image.authors="Galen Guyer <galen@galenguyer.com"

WORKDIR /go/src/genericbot
COPY . .

RUN go get -d -v ./...
RUN go build -v

FROM docker.io/library/alpine:3.14
COPY config.yaml /config.yaml
COPY --from=builder /go/src/genericbot/genericbot /genericbot
ENTRYPOINT ["/genericbot"]
