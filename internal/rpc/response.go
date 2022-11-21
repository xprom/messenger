package rpc

import "net/http"

// Response описывает структуру RPC-ответа
// В ответе должен быть один объект: либо результат в result, либо ошибка в error
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  *Error      `json:"error,omitempty"`
}

// NewResult создаёт Response с произвольным результатом
func NewResult(result interface{}) *Response {
	return &Response{Result: result}
}

// NewFromError создаёт Response из произвольной ошибки
func NewFromError(err error) *Response {
	code := http.StatusInternalServerError
	// если ошибка типа rpc.Error, из неё можно получить код
	if e, ok := err.(Error); ok {
		code = e.Code
	}

	return &Response{Error: &Error{Code: code, Message: err.Error()}}
}
