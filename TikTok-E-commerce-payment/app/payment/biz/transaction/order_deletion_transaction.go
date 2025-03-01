package transaction

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"gorm.io/gorm"
)

// DeleteOrderTransaction 删除订单事务
func DeleteOrderTransaction(ctx context.Context, db *gorm.DB, orderID string) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 删除订单
	if err := tx.Where("id = ?", orderID).Delete(&models.PaymentOrder{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
