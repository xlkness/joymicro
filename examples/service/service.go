package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/service"
)

type Shop struct {
}

func (s *Shop) Buy(ctx context.Context, req *int, resp *int) error {
	*resp = *req + 1
	fmt.Printf("buy, receive %v response %v\n", *req, *resp)
	return nil
}

func main() {
	s, err := service.New("192.168.1.188:8888", []string{"192.168.1.188:2332"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("prepare register service...\n")

	err = s.RegisterOneService("micro.shop", new(Shop), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("prepare start service...\n")

	s.Run(":8888")
}
