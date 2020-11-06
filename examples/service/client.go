package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/client"
	"joynova.com/joynova/joymicro/tracer"
	"time"
)

func main() {
	_, err := tracer.NewJaegerTracer("192.168.1.188:6831", "micro.shop")
	if err != nil {
		panic(err)
	}

	c := client.New("micro.shop1", []string{"192.168.1.188:2332"}, time.Second*10, true)
	c.SetTracer()

	r := new(int)

	ctx := context.Background()

	//ctx = context.WithValue(ctx, "key1", "2")
	err = c.Call(ctx, "Buy1", new(int), r)
	if err != nil {
		time.Sleep(time.Second * 3)
		panic(err)
	}

	fmt.Printf("buy receive:%v\n", *r)
	time.Sleep(5 * time.Second)
}
