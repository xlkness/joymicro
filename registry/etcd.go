package registry

import (
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/rpcx-etcd/client"
	"github.com/rpcxio/rpcx-etcd/serverplugin"
	xclient "github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
)

// todo etcd插件不支持注册函数服务
func GetEtcdRegistryServerPlugin(key string, serviceAddr string, etcdAddress []string) (server.Plugin, error) {
	if key == "" {
		key = "tcp"
	}

	baseDir := getBaseDir()

	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: key + "@" + serviceAddr,
		EtcdServers:    etcdAddress,
		BasePath:       baseDir,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Second * time.Duration(3),
	}
	err := r.Start()
	return r, err
}

func GetEtcdRegistryClientPlugin(service string, etcdServerAddrs []string) xclient.ServiceDiscovery {
	d, _ := client.NewEtcdDiscovery(getBaseDir(), service, etcdServerAddrs, nil)
	return d
}

func getBaseDir() string {
	if NameSpace != "" {
		if NameSpace[len(NameSpace)-1] == '/' {
			if DefaultBaseDir[0] == '/' {
				return NameSpace + DefaultBaseDir[1:]
			}
			return NameSpace + DefaultBaseDir
		}
		if DefaultBaseDir[0] == '/' {
			return NameSpace + DefaultBaseDir
		}
		return NameSpace + "/" + DefaultBaseDir
	}
	return DefaultBaseDir
}
