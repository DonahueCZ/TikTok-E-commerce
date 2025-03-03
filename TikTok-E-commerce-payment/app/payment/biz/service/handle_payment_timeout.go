package service

import (
	"context"
	"errors"
	"time"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type HandlePaymentTimeoutService struct {
	ctx context.Context
}

// NewHandlePaymentTimeoutService 创建 HandlePaymentTimeoutService
func NewHandlePaymentTimeoutService(ctx context.Context) *HandlePaymentTimeoutService {
	return &HandlePaymentTimeoutService{ctx: ctx}
}

// Run 处理支付超时逻辑
func (s *HandlePaymentTimeoutService) Run(req *payment.PaymentTimeoutRequest) (resp *payment.PaymentResponse, err error) {
	// 1. 通过订单ID查询订单
	order, err := mysql.OrderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 判断订单是否已支付
	if order.Status == "paid" {
		return &payment.PaymentResponse{
			Success: false,
			Message: "订单已支付，无法超时取消",
		}, nil
	}

	// 3. 检查订单是否超时（假设超时时间为30分钟）
	if time.Since(order.CreatedAt) < 30*time.Minute {
		return &payment.PaymentResponse{
			Success: false,
			Message: "订单未超时，无需处理",
		}, nil
	}

	// 4. 更新订单状态为 "timeout"
	err = mysql.OrderRepo.UpdateOrderStatus(s.ctx, req.OrderId, "timeout")
	if err != nil {
		return nil, err
	}

	// 5. 返回成功响应
	resp = &payment.PaymentResponse{
		Success: true,
		Message: "订单超时，已取消",
	}
	return resp, nil
}
