package hello

import (
	"context"
	apiv1 "github.com/xavierxcn/apiserver/api/v1"
	"github.com/xavierxcn/apiserver/internal/serve/pkg/errno"
)

// Service hello service
type Service struct {
}

// Hello hello handler func
func (h Service) Hello(ctx context.Context, hello *apiv1.ReqHello) (*apiv1.RspHello, error) {
	if hello.Hello != "hello" {
		return nil, errno.ErrNotFound
	}

	return &apiv1.RspHello{
		Hello: "hello",
	}, nil
}
