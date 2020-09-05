package registry

import (
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

// todo etcd插件不支持注册函数服务
func GetEtcdRegistryServerPlugin(serviceAddr string, etcdAddress []string) (server.Plugin, error) {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + serviceAddr,
		EtcdServers:    etcdAddress,
		BasePath:       DefaultBaseDir,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Second * time.Duration(60),
	}
	err := r.Start()
	return r, err
}

func GetEtcdRegistryClientPlugin(service string, etcdServerAddrs []string) client.ServiceDiscovery {
	d := client.NewEtcdDiscovery(DefaultBaseDir, service, etcdServerAddrs, nil)
	return d
}
