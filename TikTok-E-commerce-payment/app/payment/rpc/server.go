package rpc

import (
	"TikTok-E-commerce-payment/app/payment/handler"
	"TikTok-E-commerce-payment/kitex_gen/payment_proto_idl/idl/paymentservice/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
)

// NewPaymentService 创建一个新的 PaymentService 实例
func NewPaymentService() *handler.OrderHandler {
	return handler.NewOrderHandler()
}

// StartRPCServer 启动一个新的 Kitex RPC 服务
func StartRPCServer() {
	// 设置 Kitex 服务
	svr := paymentservice.NewServer(NewPaymentService())

	// 启动服务
	err := svr.Run()
	if err != nil {
		klog.Fatalf("启动 RPC 服务失败: %v", err)
	}
}
