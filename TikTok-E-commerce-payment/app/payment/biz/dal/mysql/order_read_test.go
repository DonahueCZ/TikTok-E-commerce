package mysql

import (
	"context"
	"testing"
	"time"

	"TikTok-E-commerce-payment/app/payment/biz/models"
	"github.com/stretchr/testify/assert"
)

// 测试 GetOrderByID 方法
func TestGetOrderByID(t *testing.T) {
	db := setupTestDB()
	repo := NewOrderRepository(db)

	// 1️⃣ 插入测试订单
	order := &models.PaymentOrder{
		OrderID:       "order_789",
		Amount:        75.00,
		Status:        "pending",
		PaymentMethod: "bank_transfer",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	err := db.Create(order).Error
	assert.NoError(t, err)

	// 2️⃣ 获取订单
	fetchedOrder, err := repo.GetOrderByID(context.Background(), "order_789")
	assert.NoError(t, err)
	assert.NotNil(t, fetchedOrder)
	assert.Equal(t, order.OrderID, fetchedOrder.OrderID)
	assert.Equal(t, order.Amount, fetchedOrder.Amount)
	assert.Equal(t, order.Status, fetchedOrder.Status)

	// 3️⃣ 获取一个不存在的订单
	nonExistentOrder, err := repo.GetOrderByID(context.Background(), "non_existent")
	assert.NoError(t, err)
	assert.Nil(t, nonExistentOrder)
}
