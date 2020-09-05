package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/proto/proto-gen-joymicro/proto"
	"time"
)

func main() {
	c := proto.NewShopService([]string{"192.168.1.188:2332"}, time.Second*10, true)
	resp, err := c.Buy(context.Background(), &proto.Request{
		Name: "lk",
		Num:  123,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("buy response:%+v\n", resp.String())

	resp, err = c.Buy(context.Background(), &proto.Request{
		Name: "lk",
		Num:  234,
	})
	fmt.Printf("buy response:%+v\n", resp.String())
}
