package main


import (
	"joynova.com/joynova/joymicro/service"
	"joynova.com/joynova/joymicro/examples/hello/proto"
	"joynova.com/joynova/joymicro/examples/hello/handler"
)


func main() {
	s, err := service.New(":8888", []string{"127.0.0.1:2382"})
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
