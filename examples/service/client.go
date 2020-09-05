package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/client"
	"time"
)

func main() {
	c := client.New("micro.shop", []string{"192.168.1.188:2332"}, time.Second*10, true)
	p := new(int)
	*p = 2
	r := new(int)
	err := c.Call(context.Background(), "Buy", p, r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("buy receive:%v\n", *r)

}
