package service

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
)

// GetOrderByID 通过订单 ID 查询订单
func (s *orderService) GetOrderByID(ctx context.Context, orderID string) (*models.PaymentOrder, error) {
	return s.orderRepo.GetOrderByID(ctx, orderID)
}
