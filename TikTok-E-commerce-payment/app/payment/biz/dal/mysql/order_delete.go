package mysql

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
)

// DeleteOrder 删除订单
func (r *orderRepository) DeleteOrder(ctx context.Context, orderID string) error {
	if err := r.DB.Where("order_id = ?", orderID).Delete(&models.PaymentOrder{}).Error; err != nil {
		return err
	}
	return nil
}
