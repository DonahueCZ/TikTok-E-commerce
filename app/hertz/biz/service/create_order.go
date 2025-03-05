package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateOrderService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewCreateOrderService(ctx context.Context, reqCtx *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *CreateOrderService) Run(req *rpcPayment.CreateOrderRequest) (*rpcPayment.OrderResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.CreateOrder(s.ctx, req)
	if err != nil {
		fmt.Printf("[CreateOrderService] 创建订单失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.OrderResponse{
		OrderId: resp.OrderId,
		Amount:  resp.Amount,
		Status:  resp.Status,
	}, nil
}
