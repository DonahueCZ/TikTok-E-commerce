package service

import (
	"context"
	"errors"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type HandlePaymentTimeoutService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewHandlePaymentTimeoutService 创建 HandlePaymentTimeoutService
func NewHandlePaymentTimeoutService(ctx context.Context, orderRepo mysql.OrderRepository) *HandlePaymentTimeoutService {
	return &HandlePaymentTimeoutService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 处理支付超时
func (s *HandlePaymentTimeoutService) Run(req *payment.PaymentTimeoutRequest) (*payment.PaymentResponse, error) {
	// 查询订单
	order, err := s.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 订单已支付，无法超时取消
	if order.Status == "paid" {
		return &payment.PaymentResponse{
			Status:  "failed",
			Message: "订单已支付，无法超时取消",
		}, nil
	}

	// 更新订单状态
	err = s.orderRepo.UpdateOrderStatus(req.OrderId, "timeout")
	if err != nil {
		return nil, err
	}

	return &payment.PaymentResponse{
		Status:  "timeout",
		Message: "订单已超时取消",
	}, nil
}
