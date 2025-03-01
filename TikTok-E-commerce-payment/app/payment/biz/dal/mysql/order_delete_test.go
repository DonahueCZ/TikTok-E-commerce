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

// 初始化测试数据库（使用 SQLite）
func setupDeleteTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to SQLite in-memory database")
	}
	db.AutoMigrate(&models.PaymentOrder{}) // 自动迁移 PaymentOrder 表
	return db
}

// 测试 DeleteOrder 方法
func TestDeleteOrder(t *testing.T) {
	// 初始化数据库和 repository
	db := setupDeleteTestDB()
	repo := NewOrderRepository(db)

	// 创建测试订单
	order := &models.PaymentOrder{
		OrderID:       "order_123",
		Amount:        99.99,
		Status:        "pending",
		PaymentMethod: "wechat",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := db.Create(order).Error
	assert.NoError(t, err)

	// 确保订单被成功创建
	var checkOrder models.PaymentOrder
	err = db.First(&checkOrder, "order_id = ?", "order_123").Error
	assert.NoError(t, err)
	assert.Equal(t, "order_123", checkOrder.OrderID)

	// 删除订单
	err = repo.DeleteOrder(context.Background(), "order_123")
	assert.NoError(t, err)

	// 查询订单，确保已经删除
	err = db.First(&checkOrder, "order_id = ?", "order_123").Error
	assert.Error(t, err) // 期望返回错误，因为订单应该已经删除
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
