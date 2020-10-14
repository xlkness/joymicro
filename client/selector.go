package client

import (
	"context"
	"fmt"
	"github.com/edwingeng/doublejump"
	"github.com/smallnest/rpcx/client"
	"hash/fnv"
	"math/rand"
	"sort"
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
		server := ms.servers[rand.Intn(len(ms.servers))]
		strs := strings.SplitN(server, "@", 2)
		if len(strs) != 2 {
			return fmt.Sprintf("etcd server key parse error:%v", server)
		}

		return "tcp@" + strs[1]
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

// consistentHashSelector selects based on JumpConsistentHash.
type consistentHashSelector struct {
	h       *doublejump.Hash
	servers []string
}

func NewConsistentHashSelector() client.Selector {
	h := doublejump.NewHash()
	ss := make([]string, 0)
	//for k := range servers {
	//	ss = append(ss, k)
	//	h.Add(k)
	//}

	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })
	return &consistentHashSelector{servers: ss, h: h}
}

func (s *consistentHashSelector) Select(ctx context.Context, servicePath, serviceMethod string, args interface{}) string {
	ss := s.servers
	if len(ss) == 0 {
		return ""
	}

	key := ctx.Value("select_key")

	if key == nil || key == "" {
		return s.servers[rand.Intn(len(s.servers))]
	}

	uintKey := genKey(servicePath, serviceMethod, key)
	selected, _ := s.h.Get(uintKey).(string)
	return selected
}

func (s *consistentHashSelector) UpdateServer(servers map[string]string) {
	ss := make([]string, 0, len(servers))
	for k := range servers {
		s.h.Add(k)
		ss = append(ss, k)
	}

	sort.Slice(ss, func(i, j int) bool { return ss[i] < ss[j] })

	for _, k := range s.servers {
		if servers[k] == "" { // remove
			s.h.Remove(k)
		}
	}
	s.servers = ss
}

func genKey(options ...interface{}) uint64 {
	keyString := ""
	for _, opt := range options {
		keyString = keyString + "/" + toString(opt)
	}

	return HashString(keyString)
}

func toString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

func HashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

//func Hash(key uint64, buckets int32) int32 {
//	if buckets <= 0 {
//		buckets = 1
//	}
//
//	var b, j int64
//
//	for j < int64(buckets) {
//		b = j
//		key = key*2862933555777941757 + 1
//		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
//	}
//
//	return int32(b)
//}
