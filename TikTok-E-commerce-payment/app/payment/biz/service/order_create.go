package service

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"time"
)

// CreateOrder 创建订单
func (s *orderService) CreateOrder(ctx context.Context, order *models.PaymentOrder) (*models.PaymentOrder, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Status = "pending" // 订单初始状态
	return s.orderRepo.CreateOrder(ctx, order)
}
