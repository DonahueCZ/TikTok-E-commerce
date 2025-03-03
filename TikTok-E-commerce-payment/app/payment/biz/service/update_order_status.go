package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type UpdateOrderStatusService struct {
	ctx context.Context
}

// NewUpdateOrderStatusService 创建 UpdateOrderStatusService
func NewUpdateOrderStatusService(ctx context.Context) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{ctx: ctx}
}

// Run 更新订单状态
func (s *UpdateOrderStatusService) Run(req *payment.UpdateStatusRequest) (resp *payment.PaymentResponse, err error) {
	// 1. 查询订单是否存在
	order, err := mysql.OrderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 订单状态检查，避免非法状态转换
	if order.Status == req.Status {
		return &payment.PaymentResponse{
			Success: false,
			Message: "订单状态未变化",
		}, nil
	}

	// 3. 更新订单状态
	err = mysql.OrderRepo.UpdateOrderStatus(s.ctx, req.OrderId, req.Status)
	if err != nil {
		return nil, err
	}

	// 4. 返回更新成功的响应
	resp = &payment.PaymentResponse{
		Success: true,
		Message: "订单状态更新成功",
	}
	return resp, nil
}
