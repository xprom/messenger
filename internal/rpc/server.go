package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Server описывает структуру RPC-сервера
// Хранит в себе мапу с обработчиками, фабрику валидаторов и обработчики базового функционала
type Server struct {
	handlers handlerMap
}

// NewServer создаёт новый экземпляр Server
func NewServer() *Server {
	return &Server{
		handlers: make(handlerMap),
	}
}

// Register регистрирует обработчик для метода
func (s *Server) Register(method string, handler Handler) error {
	s.handlers[method] = &handlerInfo{handler}
	return nil
}

// ServeHTTP обрабатывает все HTTP-запросы
func (s *Server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	response, err := s.HandleRequest(req)
	// при вызове обработчика возникла непредвиденная ошибка
	if err != nil {
		write(writer, NewFromError(err))
		return
	}

	write(writer, response)
}

// HandleRequest обрабатывает
func (s *Server) HandleRequest(req *http.Request) (*Response, error) {
	ctx := req.Context()

	handlerInfo, err := s.handlers.get(req.URL.Path)
	if err != nil {
		return nil, err
	}

	request := NewRequest(req)
	response, err := handlerInfo.handler.Handle(ctx, request)

	return response, err
}

func write(writer http.ResponseWriter, response *Response) {
	if response.Error != nil {
		writer.WriteHeader(response.Error.Code)
	} else {
		writer.WriteHeader(http.StatusOK)
	}

	bytes, _ := json.Marshal(response)
	_, err := writer.Write(bytes)
	if err != nil {
		fmt.Printf("[error] failed to write: %#v", err)
	}
}
