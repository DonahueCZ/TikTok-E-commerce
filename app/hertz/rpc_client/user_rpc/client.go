package userrpcclent

import (
	config "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/conf"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/user/userservice"
	userrpccl "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/user"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
)

func GetUserRpcClient() (cli userservice.Client, err error) {
	conf := config.GetConf()
	r, err := etcdRegistry.NewEtcdResolver(conf.Registry.RegistryAddress)
	if err != nil {
		return nil, err
	}
	cli, err = userservice.NewClient(
		"user",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}

func InitUserRpcClient() {
	conf := config.GetConf()
	r, err := etcdRegistry.NewEtcdResolver(conf.Registry.RegistryAddress)
	if err != nil {
		klog.Error(err)
	}
	userrpccl.InitClientWithDefaultName(client.WithResolver(r))
}
