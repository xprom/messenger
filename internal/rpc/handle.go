package rpc

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Handler описывает интерфейс для обработчика RPC-вызова, аналогичный http.Handler
type Handler interface {
	Handle(ctx context.Context, request *Request) (*Response, error)
}

// HandlerFunc описывает интерфейс инлайновой функции-обработчика RPC-вызова, аналогичный http.HandlerFunc
type HandlerFunc func(ctx context.Context, request *Request) (*Response, error)

type handlerInfo struct {
	handler Handler
}

// handlerMap хранит зарегистрированные обработчики
type handlerMap map[string]*handlerInfo

// get ищет обработчик по url запроса
func (m handlerMap) get(url string) (*handlerInfo, error) {
	method := strings.TrimSuffix(strings.TrimPrefix(url, "/"), "/")

	if handler, ok := m[method]; ok {
		return handler, nil
	}

	return nil, Error{Code: http.StatusNotFound, Message: fmt.Sprintf("method not found: %s", method)}
}
