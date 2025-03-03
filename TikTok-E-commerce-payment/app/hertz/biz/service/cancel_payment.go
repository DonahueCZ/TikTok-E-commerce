package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CancelPaymentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCancelPaymentService(Context context.Context, RequestContext *app.RequestContext) *CancelPaymentService {
	return &CancelPaymentService{RequestContext: RequestContext, Context: Context}
}

func (h *CancelPaymentService) Run(req *rpcPayment.CancelPaymentRequest) (*rpcPayment.PaymentResponse, error) {
	// 调用 RPC 层的 CancelPayment 方法
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.CancelPayment(h.Context, req)
	if err != nil {
		fmt.Println("取消支付失败:", err)
		return nil, err
	}

	// 返回处理后的数据
	resp := &rpcPayment.PaymentResponse{
		Success: rpcResp.Success,
		Message: rpcResp.Message,
	}
	return resp, nil
}
