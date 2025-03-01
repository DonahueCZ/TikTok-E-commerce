package service

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"errors"

	"TikTok-E-commerce-payment/app/payment/biz/payment_gateway" // ✅ 确保正确导入
)

// ProcessPayment 处理支付
func (s *orderService) ProcessPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(ctx, payment.OrderID)
	if err != nil || order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 获取支付方式
	factory := payment_gateway.NewPaymentFactory() // ✅ 确保正确调用
	strategy, err := factory.GetPaymentStrategy(order.PaymentMethod)
	if err != nil {
		return nil, err
	}

	// 3. 调用支付接口
	paymentReq := payment_gateway.PaymentRequest{ // ✅ 确保使用正确的 struct
		OrderID:     order.OrderID,
		Amount:      order.Amount,
		PaymentType: order.PaymentMethod,
	}

	_, err = strategy.Pay(paymentReq)
	if err != nil {
		return nil, err
	}

	// 4. 更新订单状态
	err = s.orderRepo.UpdateOrderStatus(ctx, payment.OrderID, "paid")
	if err != nil {
		return nil, err
	}

	// 5. 返回支付成功信息
	payment.Status = "paid"
	return payment, nil
}
