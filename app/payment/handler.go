package main

import (
	"context"
	payment "github.com/MelodyDeep/TikTok-E-commerce/app/payment/kitex_gen/rpc/payment"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// ProcessPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) ProcessPayment(ctx context.Context, req *payment.PaymentRequest) (resp *payment.PaymentResponse, err error) {
	resp, err = service.NewProcessPaymentService(ctx).Run(req)

	return resp, err
}

// CancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CancelPayment(ctx context.Context, req *payment.CancelRequest) (resp *payment.PaymentResponse, err error) {
	resp, err = service.NewCancelPaymentService(ctx).Run(req)

	return resp, err
}

// HandlePaymentTimeout implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) HandlePaymentTimeout(ctx context.Context, req *payment.PaymentTimeoutRequest) (resp *payment.PaymentResponse, err error) {
	resp, err = service.NewHandlePaymentTimeoutService(ctx).Run(req)

	return resp, err
}

// GetOrderByID implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetOrderByID(ctx context.Context, req *payment.OrderRequest) (resp *payment.OrderResponse, err error) {
	resp, err = service.NewGetOrderByIDService(ctx).Run(req)

	return resp, err
}

// UpdateOrderStatus implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) UpdateOrderStatus(ctx context.Context, req *payment.UpdateStatusRequest) (resp *payment.PaymentResponse, err error) {
	resp, err = service.NewUpdateOrderStatusService(ctx).Run(req)

	return resp, err
}

// DeleteOrder implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) DeleteOrder(ctx context.Context, req *payment.DeleteOrderRequest) (resp *payment.PaymentResponse, err error) {
	resp, err = service.NewDeleteOrderService(ctx).Run(req)

	return resp, err
}

// CreateOrder implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CreateOrder(ctx context.Context, req *payment.CreateOrderRequest) (resp *payment.OrderResponse, err error) {
	resp, err = service.NewCreateOrderService(ctx).Run(req)

	return resp, err
}
