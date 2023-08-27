package memory

import (
	"sync"

	pb "github.com/tn-go-course/go-search/hw-14/grpc/messenger_proto"
)

type Memory struct {
	sync.Mutex
	list []*pb.Message
}

func (m *Memory) AddMessage(message *pb.Message) {
	m.Lock()
	defer m.Unlock()
	m.list = append(m.list, message)
	return
}

func (m *Memory) Messages() []*pb.Message {
	m.Lock()
	defer m.Unlock()
	return m.list
}
