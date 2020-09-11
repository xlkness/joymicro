package service

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"joynova.com/joynova/joymicro/registry"
	"joynova.com/joynova/joymicro/util"
)

/*

 */
type Peer2Peer struct {
	PeerKey string // 点对点通信的主键
}
type Service struct {
	Service  string
	PeerInfo *Peer2Peer
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
	m := &ServicesManager{
		Addr:      addr,
		Services:  make([]*Service, 0),
		rpcserver: server.NewServer(),
	}

	// 添加etcd注册中心
	r, err := registry.GetEtcdRegistryServerPlugin(m.Addr, etcdServerAddrs)
	if err != nil {
		return nil, err
	}
	m.rpcserver.Plugins.Add(r)

	return m, nil
}

// RegisterOneService 注册一个服务
// service：服务名
// handler：回调处理
// peer2peer：是否需要点对点功能， 不需要就给nil，需要就给一个service内唯一的主键
func (m *ServicesManager) RegisterOneService(service string, handler interface{}, peer2peer *Peer2Peer) error {
	if peer2peer != nil {
		var peer2peerService string
		peer2peerService = util.GetServicePeer2Peer(service, peer2peer.PeerKey)
		err := m.checkDuplicateService(service, peer2peerService)
		if err != nil {
			return err
		}
		err = m.rpcserver.RegisterName(peer2peerService, handler, "")
		if err != nil {
			return err
		}
	} else {
		err := m.checkDuplicateService(service)
		if err != nil {
			return err
		}
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
