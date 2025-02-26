package mysql

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
)

func (r *orderRepository) CreateOrder(ctx context.Context, order *models.PaymentOrder) (*models.PaymentOrder, error) {
	if err := r.DB.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
