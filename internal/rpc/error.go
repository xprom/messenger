package rpc

// Error описывает структуру объекта RPC-ошибки
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error реализует интерфейс error
func (e Error) Error() string {
	return e.Message
}
