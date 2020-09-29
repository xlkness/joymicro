package main

import (
	"fmt"
	"joynova.com/joynova/joymicro/proto/protoc-gen-joymicro/examples/services/shop/proto1"
	"time"
)

func main() {
	to := proto1.NewShopJsonTestService([]string{"192.168.1.188:2332"}, time.Second*3, true)
	for i := 0; i < to.NumMethod(); i++ {
		fmt.Printf("fun:%v\n", to.Method(i).Name)
	}
}
