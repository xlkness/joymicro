package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/client"
	"time"
)

func main() {
	c := client.New("micro.shop", []string{"192.168.1.188:2332"}, time.Second*10, true)
	c.SetSelector(&client.PeerSelector{})
	p := new(int)
	*p = 123
	r := new(int)
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key1", "2")
	err := c.Call(ctx, "Buy", p, r)
	if err != nil {
		time.Sleep(time.Second * 3)
		panic(err)
	}

	fmt.Printf("buy receive:%v\n", *r)

}
