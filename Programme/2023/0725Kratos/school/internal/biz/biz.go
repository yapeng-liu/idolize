package biz

import (
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdAPI "go.etcd.io/etcd/client/v3"
	"school/internal/conf"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewTaskUsecase, NewDiscovery)

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	cli, err := etcdAPI.New(etcdAPI.Config{
		Endpoints: conf.Etcd.GetEndpoints(),
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(cli, etcd.Namespace(conf.Namespace))
}
