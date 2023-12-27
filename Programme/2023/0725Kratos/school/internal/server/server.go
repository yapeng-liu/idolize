package server

import (
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdAPI "go.etcd.io/etcd/client/v3"
	"school/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	client, err := etcdAPI.New(etcdAPI.Config{
		Endpoints: conf.Etcd.GetEndpoints(),
	})

	if err != nil {
		panic(err)
	}
	return etcd.New(client, etcd.Namespace(conf.Namespace))
}
