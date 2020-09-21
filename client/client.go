package client

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"joynova.com/joynova/joymicro/registry"
	"joynova.com/joynova/joymicro/util"
	"sync"
	"time"
)

type Service struct {
	ServiceName string
	etcdAddrs   []string
	// 阻塞调用超时时间
	callTimeout time.Duration
	// 默认跟rpc server不是永久连接，如果实时通讯量大的话，设置为true，
	// 字段作用：如果跟server读超时就关闭socket连接，等待之后的请求重新connect，
	// 用来避免长链接，有通信需求的双方节点形成强联通图，无用established套接字太多
	isPermanentSocketLink bool
	client                client.XClient
	peerServicesLock      *sync.Mutex
}

// New 创建对某个节点的rpc客户端管理结构
// etcdServerAddrs:etcd服务的多个节点地址
// callTimeout:调用服务超时时间
// isPermanentSocketLink:默认跟rpc server不是永久连接，如果实时通讯量大的话，设置为true，
// 						 字段作用：如果跟server读超时就关闭socket连接，等待之后的请求重新connect，
// 						 用来避免长链接，有通信需求的双方节点形成强联通图，无用established套接字太多
func New(service string, etcdServerAddrs []string, callTimeout time.Duration, isPermanentSocketLink bool) *Service {
	etcdServerAddrs = util.PreHandleEtcdHttpAddrs(etcdServerAddrs)

	conf := client.DefaultOption

	// 默认维持2min连接，读超时就和服务器断开链接
	conf.IdleTimeout = time.Minute * 2
	if isPermanentSocketLink {
		conf.Heartbeat = true
		conf.HeartbeatInterval = time.Second * 30
	}

	//conf.ReadTimeout = time.Second * 10
	//conf.WriteTimeout = time.Second * 10

	d := registry.GetEtcdRegistryClientPlugin(service, etcdServerAddrs)
	xclient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, conf)

	c := &Service{
		ServiceName:           service,
		etcdAddrs:             etcdServerAddrs,
		client:                xclient,
		callTimeout:           callTimeout,
		isPermanentSocketLink: isPermanentSocketLink,
		peerServicesLock:      &sync.Mutex{},
	}

	return c
}

func (s *Service) SetSelector(selector client.Selector) {
	s.client.SetSelector(selector)
}

// Call 根据负载算法从服务中挑一个调用
func (s *Service) Call(ctx context.Context, method string, args interface{}, reply interface{}) error {
	newCtx, f := context.WithTimeout(ctx, s.callTimeout)
	defer f()
	return s.client.Call(newCtx, method, args, reply)
}

/*
	以下为脱离服务概念的接口
*/

// CallAll 调用所有节点，有一个调用返回错误，整个调用都错误
func (s *Service) CallAll(ctx context.Context, method string, args interface{}, reply interface{}) error {
	newCtx, f := context.WithTimeout(ctx, s.callTimeout*3)
	defer f()
	return s.client.Broadcast(newCtx, method, args, reply)
}
