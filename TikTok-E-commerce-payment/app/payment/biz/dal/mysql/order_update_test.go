package mysql

import (
	"context"
	"testing"
	"time"

	"TikTok-E-commerce-payment/app/payment/biz/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 初始化测试数据库
func setupUpdateTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to SQLite in-memory database")
	}
	db.AutoMigrate(&models.PaymentOrder{}) // 自动迁移 PaymentOrder 表
	return db
}

// 测试 UpdateOrderStatus 方法
func TestUpdateOrderStatus(t *testing.T) {
	db := setupUpdateTestDB()
	repo := NewOrderRepository(db)

	// 1️⃣ 插入测试订单
	order := &models.PaymentOrder{
		OrderID:       "order_456",
		Amount:        50.00,
		Status:        "pending",
		PaymentMethod: "alipay",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := db.Create(order).Error
	assert.NoError(t, err)

	// 2️⃣ 更新订单状态
	newStatus := "completed"
	err = repo.UpdateOrderStatus(context.Background(), "order_456", newStatus)
	assert.NoError(t, err)

	// 3️⃣ 查询数据库，验证状态是否更新
	var updatedOrder models.PaymentOrder
	err = db.First(&updatedOrder, "order_id = ?", "order_456").Error
	assert.NoError(t, err)
	assert.Equal(t, newStatus, updatedOrder.Status)
}
