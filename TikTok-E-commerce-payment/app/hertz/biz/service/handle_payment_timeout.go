package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type HandlePaymentTimeoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHandlePaymentTimeoutService(Context context.Context, RequestContext *app.RequestContext) *HandlePaymentTimeoutService {
	return &HandlePaymentTimeoutService{RequestContext: RequestContext, Context: Context}
}

func (h *HandlePaymentTimeoutService) Run(req *rpcPayment.HandlePaymentTimeoutRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.HandlePaymentTimeout(h.Context, req)
	if err != nil {
		fmt.Println("处理支付超时失败:", err)
		return nil, err
	}

	resp := &rpcPayment.PaymentResponse{
		Success: rpcResp.Success,
		Message: rpcResp.Message,
	}
	return resp, nil
}
