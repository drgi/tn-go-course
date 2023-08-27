package storage

import pb "github.com/tn-go-course/go-search/hw-14/grpc/messenger_proto"

type Storage interface {
	AddMessage(*pb.Message)
	Messages() []*pb.Message
}
