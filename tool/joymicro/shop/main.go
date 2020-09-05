package main


import (
	"joynova.com/joynova/joymicro/service"
	"shop/proto"
	"shop/handler"
)


func main() {
	s, err := srevice.New(":8888", []string{"127.0.0.1:2382"})
		if err != nil {
		panic(err)
	}
	err = proto.RegisterShopHandler(s, new(ShopHandler))
	if err != nil {
		panic(err)
	}

	fmt.Printf("start service ...\n")
	err = s.Run()
	if err != nil {
		panic(err)
	}
}
