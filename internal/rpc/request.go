package rpc

import (
	"io"
	"net/http"
)

// Request описывает структуру RPC-запроса
type Request struct {
	RequestID string
	Header    http.Header
	Body      io.ReadCloser
}

// NewRequest возвращает новый Request
func NewRequest(req *http.Request) *Request {
	return &Request{
		RequestID: req.Header.Get("X-Request-Id"),
		Header:    req.Header,
		Body:      req.Body,
	}
}