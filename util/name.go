package util

import (
	"fmt"
	"strings"
)

func GetServicePeer2Peer(service, peerKey string) string {
	return fmt.Sprintf("%s-%s", service, peerKey)
}

// preHandleEtcdHttpAddrs rpcx引用的docker初始化etcd的库会将地址默认加http://
func PreHandleEtcdHttpAddrs(addrs []string) []string {
	newAddrs := make([]string, 0, len(addrs))
	for _, v := range addrs {
		idx := strings.LastIndex(v, "//")
		if idx > 0 {
			v = v[idx+2:]
		}
		newAddrs = append(newAddrs, v)
	}
	return newAddrs
}
