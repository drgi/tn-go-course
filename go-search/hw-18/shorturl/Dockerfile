FROM golang:1.21

WORKDIR /go/src/shorturl
COPY . .

RUN go mod download
WORKDIR /go/src/shorturl/cmd/server
RUN go build -o shorturl
RUN ls -l

EXPOSE 80

CMD ["./shorturl"]