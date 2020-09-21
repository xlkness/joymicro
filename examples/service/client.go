package main

import (
	"context"
	"fmt"
	"joynova.com/joynova/joymicro/client"
	"strings"
	"time"
)

type MySelector struct {
	servers []string
}

func (ms *MySelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	key := ctx.Value("key")
	if key == nil {
		return ""
	}
	if len(ms.servers) <= 0 {
		return ""
	}

	for _, server := range ms.servers {
		strs := strings.SplitN(server, "@", 2)
		if len(strs) != 2 {
			return fmt.Sprintf("server parse error:%v", server)
		}
		if strs[0] == key {
			return "tcp@" + strs[1]
		}
	}

	return ""
}

func (ms *MySelector) UpdateServer(servers map[string]string) {
	ss := make([]string, 0, len(servers))

	for k := range servers {
		ss = append(ss, k)
		delete(servers, k)
	}

	for _, v := range ss {
		// 重新修改map值，使xclient工作正确
		strs := strings.SplitN(v, "@", 2)
		if len(strs) == 2 {
			servers["tcp@"+strs[1]] = v
		}
	}

	ms.servers = ss
}

func main() {
	c := client.New("micro2/micro.shop", []string{"192.168.1.188:2332"}, time.Second*10, true)
	c.SetSelector(&MySelector{})
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
