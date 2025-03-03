package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteOrderService(Context context.Context, RequestContext *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteOrderService) Run(req *rpcPayment.DeleteOrderRequest) (*rpcPayment.PaymentResponse, error) {
	// 调用 RPC 层的 DeleteOrder 方法
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.DeleteOrder(h.Context, req)
	if err != nil {
		fmt.Println("删除订单失败:", err)
		return nil, err
	}

	// 返回处理后的数据
	resp := &rpcPayment.PaymentResponse{
		Success: rpcResp.Success,
		Message: rpcResp.Message,
	}
	return resp, nil
}
