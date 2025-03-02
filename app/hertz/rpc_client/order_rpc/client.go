package orderrpcclient

import (
	config "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/conf"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service/orderservice"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
)

func GetOrderRpcClient() (cli orderservice.Client, err error) {
	conf := config.GetConf()
	r, err := etcdRegistry.NewEtcdResolver(conf.Registry.RegistryAddress)
	if err != nil {
		return nil, err
	}
	cli, err = orderservice.NewClient(
		"order",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}

func InitOrderRpcClient() {
	conf := config.GetConf()
	r, err := etcdRegistry.NewEtcdResolver(conf.Registry.RegistryAddress)
	if err != nil {
		klog.Error(err)
	}
	orderrpc.InitClientWithDefaultName(client.WithResolver(r))
}
