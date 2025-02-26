package transaction

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"gorm.io/gorm"
)

// RefundTransaction 退款事务
func RefundTransaction(ctx context.Context, db *gorm.DB, orderID string, refundAmount float64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态为退款
	if err := tx.Model(&models.PaymentOrder{}).Where("id = ?", orderID).Update("status", "refunded").Error; err != nil {
		tx.Rollback()
		return err
	}

	// 退款金额退回到用户账户
	if err := tx.Model(&models.User{}).Where("id = ?", orderID).Update("balance", gorm.Expr("balance + ?", refundAmount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
