package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProcessPaymentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewProcessPaymentService(Context context.Context, RequestContext *app.RequestContext) *ProcessPaymentService {
	return &ProcessPaymentService{RequestContext: RequestContext, Context: Context}
}

func (h *ProcessPaymentService) Run(req *rpcPayment.ProcessPaymentRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.ProcessPayment(h.Context, req)
	if err != nil {
		fmt.Println("支付失败:", err)
		return nil, err
	}

	return &rpcPayment.PaymentResponse{
		Success: rpcResp.Success,
		Message: rpcResp.Message,
	}, nil
}
