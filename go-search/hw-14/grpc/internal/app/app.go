package app

import (
	"context"
	"math/rand"
	"time"

	"github.com/tn-go-course/go-search/hw-14/grpc/internal/storage"
	pb "github.com/tn-go-course/go-search/hw-14/grpc/messenger_proto"
)

type App struct {
	pb.UnimplementedMessangerServer

	storage storage.Storage
}

func New(s storage.Storage) *App {
	return &App{
		storage: s,
	}
}

func (a *App) Send(_ context.Context, message *pb.Message) (*pb.Empty, error) {
	message.Id = int64(rand.Int())
	message.Ts = time.Now().Unix()
	a.storage.AddMessage(message)
	return &pb.Empty{}, nil
}

func (a *App) Messages(_ *pb.Empty, stream pb.Messanger_MessagesServer) error {
	messages := a.storage.Messages()
	for _, message := range messages {
		stream.Send(message)
	}
	return nil
}
