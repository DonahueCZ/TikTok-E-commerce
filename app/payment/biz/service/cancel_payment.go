package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type CancelPaymentService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewCancelPaymentService new CancelPaymentService
func NewCancelPaymentService(ctx context.Context, orderRepo mysql.OrderRepository) *CancelPaymentService {
	return &CancelPaymentService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run handles payment cancellation
func (s *CancelPaymentService) Run(req *payment.CancelRequest) (*payment.PaymentResponse, error) {
	// 1. 查询订单是否存在
	order, err := s.orderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 如果订单已支付，不能取消
	if order.Status == "paid" {
		return nil, errors.New("订单已支付，无法取消")
	}

	// 3. 更新订单状态为 "canceled"
	err = s.orderRepo.UpdateOrderStatus(s.ctx, req.OrderId, "canceled")
	if err != nil {
		return nil, err
	}

	// 4. 返回响应
	return &payment.PaymentResponse{
		OrderId: req.OrderId,
		Status:  "canceled",
	}, nil
}
