package mysql

import (
	"TikTok-E-commerce-payment/app/payment/biz/models"
	"context"
	"gorm.io/gorm"
)

// OrderRepository 定义数据库操作接口
type OrderRepository interface {
	CreateOrder(ctx context.Context, order *models.PaymentOrder) (*models.PaymentOrder, error)
	GetOrderByID(ctx context.Context, orderID string) (*models.PaymentOrder, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) error
	DeleteOrder(ctx context.Context, orderID string) error
}

// orderRepository 结构体，封装 GORM 和 Redis 实例
type orderRepository struct {
	DB *gorm.DB
}

// NewOrderRepository 构造函数，初始化 OrderRepository 实例
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}
