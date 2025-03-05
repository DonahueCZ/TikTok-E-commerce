package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type CancelPaymentService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewCancelPaymentService 创建 CancelPaymentService
func NewCancelPaymentService(ctx context.Context, orderRepo mysql.OrderRepository) *CancelPaymentService {
	return &CancelPaymentService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 处理订单取消逻辑
func (s *CancelPaymentService) Run(req *payment.CancelRequest) (*payment.PaymentResponse, error) {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 订单已支付，不能取消
	if order.Status == "paid" {
		return nil, errors.New("订单已支付，无法取消")
	}

	// 3. 更新订单状态
	err = s.orderRepo.UpdateOrderStatus(req.OrderId, "canceled")
	if err != nil {
		return nil, err
	}

	// 4. 返回响应
	return &payment.PaymentResponse{
		Status:  "canceled",
		Message: "订单已取消",
	}, nil
}
