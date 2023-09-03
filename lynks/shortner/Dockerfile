FROM golang:1.21

WORKDIR /go/src/shortner
COPY . .
RUN ls -l
RUN go mod download
WORKDIR /go/src/shortner/cmd/server
RUN go build -o /go/bin/shortner/shortner
WORKDIR /go/bin/shortner
RUN ls -l
