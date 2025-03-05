package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateOrderStatusService struct {
	ctx    context.Context
	reqCtx *app.RequestContext
}

func NewUpdateOrderStatusService(ctx context.Context, reqCtx *app.RequestContext) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{
		ctx:    ctx,
		reqCtx: reqCtx,
	}
}

func (s *UpdateOrderStatusService) Run(req *rpcPayment.UpdateStatusRequest) (*rpcPayment.PaymentResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment")
	resp, err := rpcClient.UpdateOrderStatus(s.ctx, req)
	if err != nil {
		fmt.Printf("[UpdateOrderStatusService] 更新状态失败: %v\n", err)
		return nil, fmt.Errorf("rpc 调用失败: %w", err)
	}

	return &rpcPayment.PaymentResponse{
		Status:  resp.Status,
		Message: resp.Message,
	}, nil
}
