package mysql

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
)

// UpdateOrderStatus 更新订单状态
func (r *orderRepository) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	if err := r.DB.Model(&models.PaymentOrder{}).Where("order_id = ?", orderID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}
