package service

import (
	"TikTok-E-commerce/internal/payment/repository"
)

// OrderService 定义了订单业务逻辑的接口
type OrderService interface {
	CreateOrder(ctx context.Context, order *models.PaymentOrder) (*models.PaymentOrder, error)
	GetOrderByID(ctx context.Context, orderID string) (*models.PaymentOrder, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) error
	DeleteOrder(ctx context.Context, orderID string) error

	ProcessPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error)
	CancelPayment(ctx context.Context, paymentID string) error
	HandlePaymentTimeout(ctx context.Context, paymentID string) error
}

type orderService struct {
	orderRepo repository.OrderRepository
}

// NewOrderService 作为构造函数
func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}
