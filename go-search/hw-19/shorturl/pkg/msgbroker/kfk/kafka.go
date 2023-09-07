package kfk

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/tn-go-course/go-search/hw-19/shorturl/pkg/msgbroker"
)

type Client struct {
	Reader   *kafka.Reader
	Writer   *kafka.Writer
	handlers map[string]msgbroker.Handler
}

func New(brokers []string, topic string, groupId string) *Client {
	c := Client{}
	c.handlers = make(map[string]msgbroker.Handler)

	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	c.Writer = &kafka.Writer{
		Addr:                   kafka.TCP(brokers[0]),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &c
}

func (c *Client) Consume() {
	for {
		msg, err := c.Reader.FetchMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		handler, ok := c.handlers[msg.Topic]
		if !ok {
			log.Println("No handler for topic: ", msg.Topic)
			return
		}

		err = handler(context.Background(), msg.Value)
		if err != nil {
			log.Println("Failed to handle message: ", err)
			return
		}

		err = c.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *Client) Send(ctx context.Context, value []byte) error {
	msg := kafka.Message{
		Value: value,
	}
	return c.Writer.WriteMessages(ctx, msg)
}

func (c *Client) Register(topic string, h msgbroker.Handler) {
	c.handlers[topic] = h
}
