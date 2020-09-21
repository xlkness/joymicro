package main

import (
	"context"
	"flag"
	"fmt"
	"joynova.com/joynova/joymicro/service"
)

var no *string
var port *string

type Shop struct {
}

func (s *Shop) Buy(ctx context.Context, req *int, resp *int) error {
	*resp = *req + 1
	fmt.Printf("[%v] buy, receive %v response %v\n", *no, *req, *resp)
	return nil
}

func main() {
	no = flag.String("no", "1", "no")
	port = flag.String("p", "8888", "port")
	flag.Parse()

	s, err := service.New(*no, "192.168.1.188:"+*port, []string{"192.168.1.188:2332"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("prepare register service...\n")

	err = s.RegisterOneService("micro2/micro.shop", new(Shop))
	if err != nil {
		panic(err)
	}

	fmt.Printf("prepare start service on[%v]...\n", ":"+*port)

	err = s.Run(":" + *port)
	if err != nil {
		panic(err)
	}
}
