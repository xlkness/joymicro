package main

import (
	"joynova.com/joynova/joymicro/examples/hello/handler"
	"joynova.com/joynova/joymicro/examples/hello/proto"
	"joynova.com/joynova/joymicro/service"
)


func main() {
	s, err := service.New(":8888", []string{"192.168.1.188:2332"})
		if err != nil {
		panic(err)
	}
	err = proto.RegisterHelloHandler(s, new(handler.HelloHandler))
	if err != nil {
		panic(err)
	}

	err = s.Run()
	if err != nil {
		panic(err)
	}
}
