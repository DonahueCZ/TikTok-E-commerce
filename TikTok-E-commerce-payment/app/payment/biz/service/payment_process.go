package service

// 负责发起支付

import (
	"TikTok-E-commerce-payment/app/payment/biz/transaction"
	"context"
)

func (s *orderService) ProcessPayment(ctx context.Context, req *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	// 1. 验证订单
	if !s.ValidateOrder(req.OrderId) {
		return &payment.PaymentResponse{Status: "fail", Message: "无效订单"}, nil
	}

	// 2. 获取订单信息并验证
	order, err := s.orderRepo.GetOrderByID(ctx, req.OrderId)
	if err != nil {
		return &payment.PaymentResponse{Status: "fail", Message: "订单查询失败"}, err
	}

	if order.Status != "pending" {
		return &payment.PaymentResponse{Status: "fail", Message: "订单不可支付"}, nil
	}

	// 3. 执行支付事务
	err = transaction.ProcessPaymentTransaction(ctx, s.db, order, req)
	if err != nil {
		return &payment.PaymentResponse{Status: "fail", Message: "支付失败"}, err
	}

	return &payment.PaymentResponse{Status: "success", Message: "支付处理成功"}, nil
}
