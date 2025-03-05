package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/models"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type ProcessPaymentService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewProcessPaymentService 创建 ProcessPaymentService
func NewProcessPaymentService(ctx context.Context, orderRepo mysql.OrderRepository) *ProcessPaymentService {
	return &ProcessPaymentService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 处理支付逻辑
func (s *ProcessPaymentService) Run(req *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	// 1. 通过订单ID查询订单
	order, err := s.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 订单状态检查
	if order.Status == "paid" {
		return &payment.PaymentResponse{
			Status:  "failed",
			Message: "订单已支付，无法重复支付",
		}, nil
	}

	// 3. 模拟支付成功（这里需要你实现 `payment_gateway`）
	order.Status = "paid"

	// 4. 更新订单状态为 "paid"
	err = s.orderRepo.UpdateOrderStatus(req.OrderId, "paid")
	if err != nil {
		return nil, err
	}

	// 5. 记录支付信息
	paymentRecord := &models.Payment{
		OrderID: req.OrderId,
		Amount:  order.Amount,
		Status:  "paid",
	}
	err = s.orderRepo.CreatePaymentRecord(paymentRecord) // ⚠️ 这里要确保 `CreatePaymentRecord` 存在
	if err != nil {
		return nil, err
	}

	// 6. 返回支付成功响应
	return &payment.PaymentResponse{
		Status:  "success",
		Message: "支付成功",
	}, nil
}
