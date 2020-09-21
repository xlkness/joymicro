package client

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// PeerSelector 点对点选择器
type PeerSelector struct {
	servers   []string
	SelectFun func()
}

// Select 根据context里的select_key选择匹配的服务器进行调用
func (ms *PeerSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	if len(ms.servers) <= 0 {
		return ""
	}

	key := ctx.Value("select_key")

	if key == nil {
		return ms.servers[rand.Intn(len(ms.servers))]
		//return ""
	}

	for _, server := range ms.servers {
		strs := strings.SplitN(server, "@", 2)
		if len(strs) != 2 {
			return fmt.Sprintf("etcd server key parse error:%v", server)
		}
		if strs[0] == key {
			return "tcp@" + strs[1]
		}
	}

	return ""
}

// UpdateServer 更新服务器
func (ms *PeerSelector) UpdateServer(servers map[string]string) {
	ss := make([]string, 0, len(servers))
	servers1 := make(map[string]string, len(servers))

	for k, v := range servers {
		ss = append(ss, k)
		servers1[k] = v
		delete(servers, k)
	}

	for _, k := range ss {
		// 重新修改map值，使xclient工作正确
		strs := strings.SplitN(k, "@", 2)
		if len(strs) == 2 {
			servers["tcp@"+strs[1]] = servers1[k]
		} else {
			servers[k] = servers1[k]
		}
	}

	ms.servers = ss
}
