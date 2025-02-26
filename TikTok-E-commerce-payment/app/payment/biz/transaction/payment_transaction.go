package transaction

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"fmt"
	"gorm.io/gorm"
)

// ProcessPaymentTransaction 处理支付事务
func ProcessPaymentTransaction(ctx context.Context, db *gorm.DB, order *models.PaymentOrder, paymentDetails *models.PaymentDetails) error {
	// 开始一个新的事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 创建订单
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 扣减库存
	if err := tx.Model(&models.Product{}).Where("id = ?", paymentDetails.ProductID).Update("stock", gorm.Expr("stock - ?", paymentDetails.Quantity)).Error; err != nil {
		tx.Rollback()
		return err
	}

	//判断金额是否充足
	if err := tx.Model(&models.User{}).Where("id = ?", paymentDetails.UserID).Where("balance >= ?", paymentDetails.Amount).First(&models.User{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("余额不足")
	}

	// 扣减用户余额
	if err := tx.Model(&models.User{}).Where("id = ?", paymentDetails.UserID).Update("balance", gorm.Expr("balance - ?", paymentDetails.Amount)).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}
