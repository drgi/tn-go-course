package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tn-go-course/go-search/hw-14/grpc/internal/app"
	"github.com/tn-go-course/go-search/hw-14/grpc/internal/storage/memory"
	"github.com/tn-go-course/go-search/hw-14/grpc/messenger_proto"
	"google.golang.org/grpc"
)

const (
	port = 8000
)

func main() {
	app := app.New(&memory.Memory{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	messenger_proto.RegisterMessangerServer(grpcServer, app)
	log.Println("Listen on port: ", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
