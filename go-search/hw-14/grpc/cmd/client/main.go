package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/tn-go-course/go-search/hw-14/grpc/messenger_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = 8000
	host = "localhost"
)

func main() {
	grpcConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewMessangerClient(grpcConn)
	ctx := context.Background()
	err = sendMessages(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	err = writeRecivedMessages(ctx, client, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessages(ctx context.Context, client pb.MessangerClient) error {
	messages := []*pb.Message{
		{
			Text: "Привет!",
		},
		{
			Text: "Как дела?",
		},
	}

	for _, m := range messages {
		_, err := client.Send(ctx, m)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeRecivedMessages(ctx context.Context, client pb.MessangerClient, output io.Writer) error {
	stream, err := client.Messages(ctx, &pb.Empty{})
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			m, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			date := time.Unix(m.Ts, 0)
			line := fmt.Sprintf("Получено сообщение! [%s] ИД: %d, Текст: %s", date.Format(time.Kitchen), m.Id, m.Text)
			_, err = output.Write([]byte(line + "\n"))
			if err != nil {
				return err
			}
		}
	}

}
