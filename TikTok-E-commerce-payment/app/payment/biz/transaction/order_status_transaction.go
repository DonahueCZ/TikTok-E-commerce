package transaction

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"gorm.io/gorm"
)

// UpdateOrderStatusTransaction 更新订单状态事务
func UpdateOrderStatusTransaction(ctx context.Context, db *gorm.DB, orderID string, status string) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态
	if err := tx.Model(&models.PaymentOrder{}).Where("id = ?", orderID).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
