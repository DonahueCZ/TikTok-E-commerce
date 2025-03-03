package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateOrderStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderStatusService) Run(req *rpcPayment.UpdateOrderStatusRequest) (resp *rpcPayment.PaymentResponse, err error) {
	// 调用 RPC 层的 UpdateOrderStatus 方法
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.UpdateOrderStatus(h.Context, req)
	if err != nil {
		fmt.Println("更新订单状态失败:", err)
		return nil, err
	}

	// 返回处理后的数据
	resp = &rpcPayment.PaymentResponse{
		Success: rpcResp.Success,
		Message: rpcResp.Message,
	}
	return resp, nil
}
