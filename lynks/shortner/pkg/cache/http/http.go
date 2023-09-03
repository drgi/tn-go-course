package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	host       string
}

func New(host string) *Client {
	c := &Client{
		host: host,
	}
	c.httpClient = &http.Client{Timeout: 5 * time.Second}
	return c
}

func (c *Client) SetString(ctx context.Context, key string, value string) error {
	payload := &Request{
		Key:   key,
		Value: value,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		c.host,
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	response := &ResponseEnvelope{}
	b, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return err
	}
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (c *Client) GetString(ctx context.Context, key string) (string, error) {
	requestUrl := fmt.Sprintf("%s/%s", c.host, key)
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet,
		requestUrl,
		nil)
	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	response := &ResponseEnvelope{}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(b, &response)
	if err != nil {
		return "", err
	}
	if response.Error != nil {
		return response.Result, response.Error
	}
	return response.Result, nil
}
