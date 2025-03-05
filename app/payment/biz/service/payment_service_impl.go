package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

// PaymentServiceImpl 实现 Kitex 生成的服务接口
type PaymentServiceImpl struct {
	orderRepo mysql.OrderRepository // ✅ 增加 OrderRepository 作为字段
}

// NewPaymentServiceImpl 创建 PaymentServiceImpl 实例
func NewPaymentServiceImpl(orderRepo mysql.OrderRepository) *PaymentServiceImpl {
	return &PaymentServiceImpl{orderRepo: orderRepo}
}

// ProcessPayment 处理支付
func (s *PaymentServiceImpl) ProcessPayment(ctx context.Context, req *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	return NewProcessPaymentService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// CancelPayment 取消支付
func (s *PaymentServiceImpl) CancelPayment(ctx context.Context, req *payment.CancelRequest) (*payment.PaymentResponse, error) {
	return NewCancelPaymentService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// HandlePaymentTimeout 处理支付超时
func (s *PaymentServiceImpl) HandlePaymentTimeout(ctx context.Context, req *payment.PaymentTimeoutRequest) (*payment.PaymentResponse, error) {
	return NewHandlePaymentTimeoutService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// GetOrderByID 查询订单
func (s *PaymentServiceImpl) GetOrderByID(ctx context.Context, req *payment.OrderRequest) (*payment.OrderResponse, error) {
	return NewGetOrderByIDService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// UpdateOrderStatus 更新订单状态
func (s *PaymentServiceImpl) UpdateOrderStatus(ctx context.Context, req *payment.UpdateStatusRequest) (*payment.PaymentResponse, error) {
	return NewUpdateOrderStatusService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// DeleteOrder 删除订单
func (s *PaymentServiceImpl) DeleteOrder(ctx context.Context, req *payment.DeleteOrderRequest) (*payment.PaymentResponse, error) {
	return NewDeleteOrderService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}

// CreateOrder 创建订单
func (s *PaymentServiceImpl) CreateOrder(ctx context.Context, req *payment.CreateOrderRequest) (*payment.OrderResponse, error) {
	return NewCreateOrderService(ctx, s.orderRepo).Run(req) // ✅ 传入 `s.orderRepo`
}
