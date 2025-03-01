package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"TikTok-E-commerce-payment/app/payment/biz/dal/mysql"
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"TikTok-E-commerce-payment/app/payment/biz/payment_gateway"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ✅ Mock 支付策略
type MockPaymentStrategy struct{}

// ✅ Mock 支付方法
func (m *MockPaymentStrategy) Pay(req payment_gateway.PaymentRequest) (string, error) {
	if req.OrderID == "fail_order" {
		return "", errors.New("支付失败")
	}
	return "Mock Payment Success", nil
}

// ✅ Mock PaymentFactory（Mock `GetPaymentStrategy`）
type MockPaymentFactory struct{}

// ✅ `MockPaymentFactory` 实现 `GetPaymentStrategy`
func (f *MockPaymentFactory) GetPaymentStrategy(paymentType string) (payment_gateway.PaymentStrategy, error) {
	return &MockPaymentStrategy{}, nil
}

// ✅ 重新初始化 SQLite 测试数据库
func setupPaymentTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("无法连接 SQLite 内存数据库")
	}
	db.AutoMigrate(&models.PaymentOrder{}, &models.Payment{})
	return db
}

// ✅ 更新 `TestProcessPayment`
func TestProcessPayment(t *testing.T) {
	db := setupTestDB()
	repo := mysql.NewOrderRepository(db)
	service := NewOrderService(repo) // ✅ 确保 `service` 被使用

	// ✅ 插入订单
	order := &models.PaymentOrder{
		OrderID:       "order_1001",
		Amount:        100.00,
		Status:        "pending",
		PaymentMethod: "alipay",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(order)

	// ✅ 直接使用 MockPaymentFactory
	factory := &MockPaymentFactory{}

	// ✅ 运行测试
	payment := &models.Payment{ // ✅ 确保 `payment` 被使用
		PaymentID: "pay_1001",
		OrderID:   "order_1001",
		Amount:    100.00,
		Status:    "pending",
	}
	strategy, err := factory.GetPaymentStrategy(order.PaymentMethod)
	assert.NoError(t, err)

	_, err = strategy.Pay(payment_gateway.PaymentRequest{
		OrderID:     order.OrderID,
		Amount:      order.Amount,
		PaymentType: order.PaymentMethod,
	})
	assert.NoError(t, err)

	// ✅ 确保 `service.ProcessPayment()` 调用了 `payment`
	paidPayment, err := service.ProcessPayment(context.Background(), payment)
	assert.NoError(t, err)
	assert.NotNil(t, paidPayment)
	assert.Equal(t, "paid", paidPayment.Status)

	// ✅ 确保订单状态已更新
	var updatedOrder models.PaymentOrder
	db.First(&updatedOrder, "order_id = ?", "order_1001")
	assert.Equal(t, "paid", updatedOrder.Status)
}
