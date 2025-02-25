package rpc

import (
	"sync"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/conf"
	"github.com/MelodyDeep/TikTok-E-commerce/common/clientsuite"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	once          sync.Once
	commonSuite   client.Option
	ProductClient productcatalogservice.Client
)

func InitClient() {
	once.Do(func() {
		commonSuite = client.WithSuite(
			clientsuite.CommonGrpcClientSuite{
				RegistryAddr:       conf.GetConf().Hertz.RegistryAddr,
				CurrentServiceName: conf.GetConf().Hertz.Service,
			},
		)
		initProductClient()
	})
}

func initProductClient() {
	ProductClient = productcatalogservice.MustNewClient("product", commonSuite)
}
