package main

import (
	"context"
	"joynova.com/joynova/joymicro/examples/hello1/proto"
	"joynova.com/joynova/joymicro/service"
)

type Handler struct {
}

func (h *Handler) Hello(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.ErrCode = 1
	res.Msg = "hello " + req.GetName()
	return nil
}

func main() {
	s, err := service.NewWithKey("key1", "192.168.1.2:8888", []string{"192.168.1.2:2383"})
	if err != nil {
		panic(err)
	}
	err = proto.RegisterHelloHandler(s, new(Handler))
	if err != nil {
		panic(err)
	}

	err = s.Run(":8888")
	if err != nil {
		panic(err)
	}
}
