package rpcclient

import (
	"order/order_hertz_api/config"
	"order/rpc/order/kitex_gen/demo/order_service/orderservice"

	"github.com/cloudwego/kitex/client"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
)

func GetOrderRpcClient() (cli orderservice.Client, err error) {
	conf := config.GetConfig()
	r, err := etcdRegistry.NewEtcdResolver([]string{conf.EtcdHost})
	if err != nil {
		return nil, err
	}
	cli, err = orderservice.NewClient(
		conf.OrderServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
