FROM golang:1.21

WORKDIR /go/src/memcache
COPY . .
RUN ls -l
RUN go mod download
WORKDIR /go/src/memcache/cmd/server
RUN go build -o /go/bin/memcache/memcache
WORKDIR /go/bin/memcache
RUN ls -l
