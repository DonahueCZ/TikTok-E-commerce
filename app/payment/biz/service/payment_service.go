package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

// PaymentService 处理支付逻辑
type PaymentService struct{}

// ProcessPayment 处理支付请求
func (s *PaymentService) ProcessPayment(ctx context.Context, req *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	return &payment.PaymentResponse{
		Status:  "success",
		Message: "支付成功",
	}, nil
}
