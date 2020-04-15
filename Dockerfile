FROM golang:latest AS build
WORKDIR /go/src
COPY . /go/src
RUN go build ./cmd/archiver

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /root/
COPY --from=build /go/src/archiver /usr/local/bin
ENTRYPOINT ["/usr/local/bin/archiver"]
