package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type HandlePaymentTimeoutService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewHandlePaymentTimeoutService(ctx context.Context, reqCtx *app.RequestContext) *HandlePaymentTimeoutService {
	return &HandlePaymentTimeoutService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *HandlePaymentTimeoutService) Run(req *rpcPayment.PaymentTimeoutRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.HandlePaymentTimeout(s.ctx, req)
	if err != nil {
		fmt.Printf("[HandlePaymentTimeoutService] 处理超时失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.PaymentResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
