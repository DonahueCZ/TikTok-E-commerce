package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateOrderService(Context context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateOrderService) Run(req *rpcPayment.CreateOrderRequest) (*rpcPayment.OrderResponse, error) {
	// 调用 RPC 层的 CreateOrder 方法
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.CreateOrder(h.Context, req)
	if err != nil {
		fmt.Println("创建订单失败:", err)
		return nil, err
	}

	// 返回处理后的数据
	resp := &rpcPayment.OrderResponse{
		OrderId: rpcResp.OrderId,
		Status:  rpcResp.Status,
	}
	return resp, nil
}
