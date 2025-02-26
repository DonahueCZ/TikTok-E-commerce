package service

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
)

func (s *orderService) CreateOrder(ctx context.Context, order *models.PaymentOrder) (*models.PaymentOrder, error) {
	return s.orderRepo.CreateOrder(ctx, order)
}
