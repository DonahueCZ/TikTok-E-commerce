package mysql

import (
	"database/sql"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// **创建订单**
func (r *OrderRepository) CreateOrder(order *models.PaymentOrder) (string, error) {
	return order.OrderID, errors.New("CreateOrder not implemented")
}

// **查询订单**
func (r *OrderRepository) GetOrderByID(orderID string) (*models.PaymentOrder, error) {
	return nil, errors.New("GetOrderByID not implemented")
}

// **更新订单状态**
func (r *OrderRepository) UpdateOrderStatus(orderID, status string) error {
	return errors.New("UpdateOrderStatus not implemented")
}

// **新增删除订单方法**
func (r *OrderRepository) DeleteOrder(orderID string) error {
	return errors.New("DeleteOrder not implemented")
}

// **新增支付记录**
func (r *OrderRepository) CreatePaymentRecord(payment *models.Payment) error {
	return errors.New("CreatePaymentRecord not implemented")
}
