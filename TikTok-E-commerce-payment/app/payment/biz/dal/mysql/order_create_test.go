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
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to SQLite in-memory database")
	}

	// 自动迁移 PaymentOrder 表
	db.AutoMigrate(&models.PaymentOrder{})
	return db
}

// 测试 CreateOrder 方法
func TestCreateOrder(t *testing.T) {
	// 初始化数据库和 repository
	db := setupTestDB()
	repo := NewOrderRepository(db)

	// 创建测试订单数据
	order := &models.PaymentOrder{
		OrderID:       "123456",
		Amount:        100.50,
		Status:        "pending",
		PaymentMethod: "credit_card",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 执行 CreateOrder 方法
	createdOrder, err := repo.CreateOrder(context.Background(), order)

	// 断言（Assertion）
	assert.NoError(t, err)                               // 期待无错误
	assert.NotNil(t, createdOrder)                       // 期待返回非空
	assert.Equal(t, order.OrderID, createdOrder.OrderID) // 期待订单 ID 一致

	// 从数据库查询验证是否插入成功
	var fetchedOrder models.PaymentOrder
	err = db.First(&fetchedOrder, "order_id = ?", order.OrderID).Error
	assert.NoError(t, err)
	assert.Equal(t, order.OrderID, fetchedOrder.OrderID)
}
