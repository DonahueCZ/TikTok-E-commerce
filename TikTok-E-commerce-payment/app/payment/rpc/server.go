package rpc

import (
	"TikTok-E-commerce-payment/app/payment/handler"
	"TikTok-E-commerce-payment/kitex_gen/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"net"
)

// ✅ 启动 Kitex RPC 服务器
func StartRPCServer() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8000") // 监听 8000 端口

	svr := paymentservice.NewServer(
		handler.NewOrderHandler(), // 绑定 Handler
		server.WithServiceAddr(addr),
	)

	klog.Infof("RPC 服务器启动: %v", addr)

	if err := svr.Run(); err != nil {
		klog.Fatalf("RPC 服务器启动失败: %v", err)
	}
}
