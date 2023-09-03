package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"time"
)

type Client struct {
	httpClient *http.Client
	basePath   string
}

func New(baseUrl string) *Client {
	c := &Client{
		basePath: baseUrl,
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
		path.Join(c.basePath, "/storage"),
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
	return response.Error
}

func (c *Client) GetString(ctx context.Context, key string) (string, error) {
	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet,
		path.Join(c.basePath, "/storage"),
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

	return response.Result, response.Error
}
