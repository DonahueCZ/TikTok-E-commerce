package service

import (
	"context"
	"testing"
	"time"

	"TikTok-E-commerce-payment/app/payment/biz/dal/mysql"
	"TikTok-E-commerce-payment/app/payment/biz/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 初始化 SQLite 测试数据库
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("无法连接 SQLite 内存数据库")
	}
	db.AutoMigrate(&models.PaymentOrder{}) // 自动迁移 PaymentOrder 表
	return db
}

// 测试 CreateOrder 方法
func TestCreateOrder(t *testing.T) {
	db := setupTestDB()
	repo := mysql.NewOrderRepository(db)
	service := NewOrderService(repo)

	order := &models.PaymentOrder{
		OrderID:       "order_001",
		Amount:        99.99,
		Status:        "pending",
		PaymentMethod: "alipay",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	createdOrder, err := service.CreateOrder(context.Background(), order)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, createdOrder)
	assert.Equal(t, order.OrderID, createdOrder.OrderID)

	// 验证数据库是否真的有数据
	var dbOrder models.PaymentOrder
	err = db.First(&dbOrder, "order_id = ?", "order_001").Error
	assert.NoError(t, err)
	assert.Equal(t, order.OrderID, dbOrder.OrderID)
}

// 测试 GetOrderByID 方法
func TestGetOrderByID(t *testing.T) {
	db := setupTestDB()
	repo := mysql.NewOrderRepository(db)
	service := NewOrderService(repo)

	// 先插入一个订单
	order := &models.PaymentOrder{
		OrderID:       "order_002",
		Amount:        50.00,
		Status:        "pending",
		PaymentMethod: "wechat",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(order)

	// 读取订单
	fetchedOrder, err := service.GetOrderByID(context.Background(), "order_002")

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, fetchedOrder)
	assert.Equal(t, order.OrderID, fetchedOrder.OrderID)
	assert.Equal(t, order.Amount, fetchedOrder.Amount)
	assert.Equal(t, order.Status, fetchedOrder.Status)
}

// 测试 UpdateOrderStatus 方法
func TestUpdateOrderStatus(t *testing.T) {
	db := setupTestDB()
	repo := mysql.NewOrderRepository(db)
	service := NewOrderService(repo)

	// 插入一个订单
	order := &models.PaymentOrder{
		OrderID:       "order_003",
		Amount:        70.00,
		Status:        "pending",
		PaymentMethod: "bank_transfer",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(order)

	// 更新订单状态
	err := service.UpdateOrderStatus(context.Background(), "order_003", "completed")
	assert.NoError(t, err)

	// 查询数据库，验证状态是否更新
	var updatedOrder models.PaymentOrder
	err = db.First(&updatedOrder, "order_id = ?", "order_003").Error
	assert.NoError(t, err)
	assert.Equal(t, "completed", updatedOrder.Status)
}

// 测试 DeleteOrder 方法
func TestDeleteOrder(t *testing.T) {
	db := setupTestDB()
	repo := mysql.NewOrderRepository(db)
	service := NewOrderService(repo)

	// 插入一个订单
	order := &models.PaymentOrder{
		OrderID:       "order_004",
		Amount:        120.00,
		Status:        "pending",
		PaymentMethod: "credit_card",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(order)

	// 删除订单
	err := service.DeleteOrder(context.Background(), "order_004")
	assert.NoError(t, err)

	// 确保订单被删除
	var checkOrder models.PaymentOrder
	err = db.First(&checkOrder, "order_id = ?", "order_004").Error
	assert.Error(t, err) // 期望返回错误，因为订单已经删除
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
