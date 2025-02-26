package service

import (
	"TikTok-E-commerce/internal/payment/models"
	"context"
)

func (s *orderService) GetOrderByID(ctx context.Context, orderID string) (*models.PaymentOrder, error) {
	return s.orderRepo.GetOrderByID(ctx, orderID)
}
