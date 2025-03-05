package rpc_client

import (
	"fmt"

	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
)

var PaymentClient paymentservice.Client

func InitPaymentClient() {
	var err error
	PaymentClient, err = paymentservice.NewClient("payment", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		fmt.Println("支付 RPC 客户端初始化失败:", err)
	}
}
