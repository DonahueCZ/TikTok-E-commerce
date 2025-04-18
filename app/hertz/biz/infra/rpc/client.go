package rpc

import (
	"sync"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/conf"
	"github.com/MelodyDeep/TikTok-E-commerce/common/clientsuite"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	// consul "github.com/kitex-contrib/registry-consul"
)

var (
	once          sync.Once
	commonSuite   client.Option
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
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

func initCartClient(){
	CartClient = cartservice.MustNewClient("cart", commonSuite)
}
