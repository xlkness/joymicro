package registry

import (
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"time"
)

// todo etcd插件不支持注册函数服务
func GetEtcdRegistryServerPlugin(key string, serviceAddr string, etcdAddress []string) (server.Plugin, error) {
	if key == "" {
		key = "tcp"
	}
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: key + "@" + serviceAddr,
		EtcdServers:    etcdAddress,
		BasePath:       DefaultBaseDir,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Second * time.Duration(3),
	}
	err := r.Start()
	return r, err
}

func GetEtcdRegistryClientPlugin(service string, etcdServerAddrs []string) client.ServiceDiscovery {
	d := client.NewEtcdDiscovery(DefaultBaseDir, service, etcdServerAddrs, nil)
	return d
}
