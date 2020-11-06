package service

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"joynova.com/joynova/joymicro/registry"
	"joynova.com/joynova/joymicro/tracer"
	"joynova.com/joynova/joymicro/util"
)

/*

 */
type Service struct {
	Service string
}

type ServicesManager struct {
	Addr      string         // 节点提供rpc服务的地址
	Services  []*Service     // 注册的服务信息
	rpcserver *server.Server // rpc服务器，用来注册服务，服务发现等
}

// New 创建一个服务
// service:服务器名称
// addr:当前节点服务对外可以访问的地址，不是监听地址，必须为"ip+:+port"格式
// etcdServerAddrs:etcd服务地址
func New(addr string, etcdServerAddrs []string) (*ServicesManager, error) {
	etcdServerAddrs = util.PreHandleEtcdHttpAddrs(etcdServerAddrs)
	m := newServersManager(addr)

	// 添加etcd注册中心
	r, err := registry.GetEtcdRegistryServerPlugin("", m.Addr, etcdServerAddrs)
	if err != nil {
		return nil, err
	}
	m.rpcserver.Plugins.Add(r)

	return m, nil
}

// NewWithKey 创建一个带主键的服务，用于点对点通信
// service:服务器名称
// addr:当前节点服务对外可以访问的地址，不是监听地址，必须为"ip+:+port"格式
// etcdServerAddrs:etcd服务地址
func NewWithKey(key string, addr string, etcdServerAddrs []string) (*ServicesManager, error) {
	etcdServerAddrs = util.PreHandleEtcdHttpAddrs(etcdServerAddrs)
	m := newServersManager(addr)

	// 添加etcd注册中心
	r, err := registry.GetEtcdRegistryServerPlugin(key, m.Addr, etcdServerAddrs)
	if err != nil {
		return nil, err
	}
	m.rpcserver.Plugins.Add(r)

	return m, nil
}

func (m *ServicesManager) SetTracer() {
	m.rpcserver.Plugins.Add(&tracer.JaegerOpenTracingServerPlugin{})
}

// RegisterOneService 注册一个服务
// service：服务名
// handler：回调处理
func (m *ServicesManager) RegisterOneService(service string, handler interface{}) error {
	err := m.checkDuplicateService(service)
	if err != nil {
		return err
	}

	return m.rpcserver.RegisterName(service, handler, "")
}

func (m *ServicesManager) checkDuplicateService(services ...string) error {
	for _, v := range m.Services {
		for _, v1 := range services {
			if v.Service == v1 {
				return fmt.Errorf("duplicate register service:%v", v1)
			}
		}
	}

	return nil
}

// Run 启动rpc服务
// addr：监听地址，可以忽略ip，例如":8888"格式
// 注意：register过程必须在start之前
func (m *ServicesManager) Run(addr string) error {
	return m.rpcserver.Serve("tcp", addr)
}

func newServersManager(addr string) *ServicesManager {
	m := &ServicesManager{
		Addr:      addr,
		Services:  make([]*Service, 0),
		rpcserver: server.NewServer(),
	}

	//m.rpcserver.Plugins.Add(&tracer.JaegerOpenTracingServerPlugin{})

	return m
}
