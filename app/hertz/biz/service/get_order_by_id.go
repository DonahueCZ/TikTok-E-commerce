package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetOrderByIDService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewGetOrderByIDService(ctx context.Context, reqCtx *app.RequestContext) *GetOrderByIDService {
	return &GetOrderByIDService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *GetOrderByIDService) Run(req *rpcPayment.OrderRequest) (*rpcPayment.OrderResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.GetOrderByID(s.ctx, req)
	if err != nil {
		fmt.Printf("[GetOrderByIDService] 查询订单失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.OrderResponse{
		OrderId: resp.OrderId,
		Amount:  resp.Amount,
		Status:  resp.Status,
	}, nil
}
