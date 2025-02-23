package rpc

import (
	"sync"

	"github.com/MelodyDeep/TikTok-E-commerce/app/cart/conf"
	cartutils "github.com/MelodyDeep/TikTok-E-commerce/app/cart/utils"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
)

func InitClient() {
	once.Do(func() {
		//registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		//serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().RegistryAddress[0])
	cartutils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
