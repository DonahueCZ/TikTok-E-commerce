package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type ProcessPaymentService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewProcessPaymentService(ctx context.Context, reqCtx *app.RequestContext) *ProcessPaymentService {
	return &ProcessPaymentService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *ProcessPaymentService) Run(req *rpcPayment.PaymentRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.ProcessPayment(s.ctx, req)
	if err != nil {
		fmt.Printf("[ProcessPaymentService] 支付失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.PaymentResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
