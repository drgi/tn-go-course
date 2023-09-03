package redis

import (
	"context"
	"fmt"
	"time"

	r "github.com/go-redis/redis/v8"
)

type Client struct {
	c                 *r.Client
	valueLifeTimeHour int
}

func New(host, port, password string, valueLifeTimeHour int) *Client {
	redisClient := r.NewClient(&r.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       0,
	})
	return &Client{c: redisClient}
}

func (c *Client) SetString(ctx context.Context, key string, value string) error {
	return c.c.Set(ctx, key, value, time.Hour*time.Duration(c.valueLifeTimeHour)).Err()
}

func (c *Client) GetString(ctx context.Context, key string) (string, error) {
	cmd := c.c.Get(context.Background(), key)
	return cmd.Result()
}
