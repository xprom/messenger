package messenger_test

import (
	"context"

	"messenger/internal/repository"
	"messenger/internal/rpc"
)

type Handler struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) *Handler {
	return &Handler{repository}
}

func (h *Handler) Handle(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	//result, err := h.repository.List(ctx)
	//
	//if err != nil {
	//	return rpc.NewFromError(err), err
	//}

	return rpc.NewResult(Response{
		Success: true,
		Message: "привет",
		AdditionalData: ResponseInner{
			AdditionalInfo: "123",
		},
	}), nil
}
