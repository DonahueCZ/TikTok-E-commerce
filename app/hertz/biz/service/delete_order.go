package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteOrderService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewDeleteOrderService(ctx context.Context, reqCtx *app.RequestContext) *DeleteOrderService {
	return &DeleteOrderService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *DeleteOrderService) Run(req *rpcPayment.DeleteOrderRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.DeleteOrder(s.ctx, req)
	if err != nil {
		fmt.Printf("[DeleteOrderService] 删除订单失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.PaymentResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
