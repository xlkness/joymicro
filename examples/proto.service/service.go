package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/proto/proto-gen-joymicro/proto"
	"joynova.com/joynova/joymicro/service"
)

type ShopHandler struct {
}

func (s *ShopHandler) Buy(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.ErrCode = 1
	res.Num = req.Num
	res.CurMoney = 345
	fmt.Printf("buy receive %+v response %+v\n", req, res)
	return nil
}

func main() {
	s, err := service.New(":8888", []string{"192.168.1.188:2332"})
	if err != nil {
		panic(err)
	}
	proto.RegisterShopHandler(s, new(ShopHandler))
	s.Run()
}
