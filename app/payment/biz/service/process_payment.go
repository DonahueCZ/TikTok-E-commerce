package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/models"

	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type ProcessPaymentService struct {
	ctx context.Context
}

// NewProcessPaymentService 创建 ProcessPaymentService
func NewProcessPaymentService(ctx context.Context) *ProcessPaymentService {
	return &ProcessPaymentService{ctx: ctx}
}

// Run 处理支付逻辑
func (s *ProcessPaymentService) Run(req *payment.PaymentRequest) (resp *payment.PaymentResponse, err error) {
	// 1. 通过订单ID查询订单
	order, err := mysql.OrderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 订单状态检查
	if order.Status == "paid" {
		return &payment.PaymentResponse{
			Success: false,
			Message: "订单已支付，无法重复支付",
		}, nil
	}

	// 3. 获取支付方式并选择对应的支付策略
	factory := payment_gateway.NewPaymentFactory()
	strategy, err := factory.GetPaymentStrategy(order.PaymentMethod)
	if err != nil {
		return nil, err
	}

	// 4. 调用支付接口进行支付
	paymentReq := payment_gateway.PaymentRequest{
		OrderID:     order.OrderID,
		Amount:      order.Amount,
		PaymentType: order.PaymentMethod,
	}
	_, err = strategy.Pay(paymentReq)
	if err != nil {
		return nil, err
	}

	// 5. 更新订单状态为 "paid"
	err = mysql.OrderRepo.UpdateOrderStatus(s.ctx, req.OrderId, "paid")
	if err != nil {
		return nil, err
	}

	// 6. 记录支付信息
	paymentRecord := &models.Payment{
		OrderID: req.OrderId,
		Amount:  order.Amount,
		Status:  "paid",
	}
	err = mysql.PaymentRepo.CreatePaymentRecord(s.ctx, paymentRecord)
	if err != nil {
		return nil, err
	}

	// 7. 返回支付成功响应
	resp = &payment.PaymentResponse{
		Success: true,
		Message: "支付成功",
	}
	return resp, nil
}
