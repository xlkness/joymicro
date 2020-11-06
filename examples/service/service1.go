package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/client"
	"joynova.com/joynova/joymicro/service"
	"joynova.com/joynova/joymicro/tracer"
	"time"
)

var (
	c1 *client.Service
	c2 *client.Service
	c3 *client.Service
)

type Shop1 struct {
}

func (s *Shop1) Buy1(ctx context.Context, req *int, resp *int) error {
	*resp = *req + 1
	fmt.Printf("1 buy, receive %v response %v\n", *req, *resp)

	err := c2.Call(ctx, "Buy2", new(int), new(int))
	if err != nil {
		return err
	}

	err = c3.Call(ctx, "Buy3", new(int), new(int))
	if err != nil {
		return err
	}

	return nil
}

type Shop2 struct {
}

func (s *Shop2) Buy2(ctx context.Context, req *int, resp *int) error {
	*resp = *req + 1
	fmt.Printf("2 buy, receive %v response %v\n", *req, *resp)

	return nil
}

type Shop3 struct {
}

func (s *Shop3) Buy3(ctx context.Context, req *int, resp *int) error {
	*resp = *req + 1
	fmt.Printf("3 buy, receive %v response %v\n", *req, *resp)
	return nil
}

func must(a interface{}, b error) interface{} {
	if b != nil {
		panic(b)
	}

	return a
}

func must1(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	must(tracer.NewJaegerTracer("192.168.1.188:6831", "micro.shop"))

	s1 := must(service.New("192.168.1.188:8081", []string{"192.168.1.188:2332"})).(*service.ServicesManager)

	s1.SetTracer()

	must1(s1.RegisterOneService("micro.shop1", new(Shop1)))
	must1(s1.RegisterOneService("micro.shop2", new(Shop2)))
	must1(s1.RegisterOneService("micro.shop3", new(Shop3)))

	go func() { must1(s1.Run(":8081")) }()

	time.Sleep(time.Second * 2)

	//c1 = client.New("micro.shop1", []string{"192.168.1.188:2332"}, time.Second*10, true)
	c2 = client.New("micro.shop2", []string{"192.168.1.188:2332"}, time.Second*10, true)
	c3 = client.New("micro.shop3", []string{"192.168.1.188:2332"}, time.Second*10, true)
	//c1.SetTracer()
	c2.SetTracer()
	c3.SetTracer()

	fmt.Printf("start ok.\n")
	select {}
}
