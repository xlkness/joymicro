package handler


import (
	"context"
	"joynova.com/joynova/joymicro/examples/hello/proto"
)

type HelloHandler struct {
}

func (s *HelloHandler) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	return nil
}
