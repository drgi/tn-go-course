package http

import "fmt"

type ResponseEnvelope struct {
	Result string    `json:"data"`
	Error  *ApiError `json:"error"`
}

type ApiError struct {
	RequestId int    `json:"requestId"`
	Message   string `json:"message"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("requestIdL: %d, message: %s", e.RequestId, e.Message)
}

type Request struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
