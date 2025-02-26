package mysql

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

// GetOrderByID 通过订单ID查询订单
func (r *orderRepository) GetOrderByID(ctx context.Context, orderID string) (*models.PaymentOrder, error) {
	var order models.PaymentOrder
	if err := r.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &order, nil
}
