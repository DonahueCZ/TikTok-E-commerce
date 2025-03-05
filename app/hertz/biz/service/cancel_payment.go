package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CancelPaymentService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewCancelPaymentService(ctx context.Context, reqCtx *app.RequestContext) *CancelPaymentService {
	return &CancelPaymentService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *CancelPaymentService) Run(req *rpcPayment.CancelRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.CancelPayment(s.ctx, req)
	if err != nil {
		fmt.Printf("[CancelPaymentService] 取消支付失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.PaymentResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
