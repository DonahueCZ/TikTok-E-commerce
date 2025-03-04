package service

import (
	"context"
	"errors"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/models"
	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
	"time"
)

type CreateOrderService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context, orderRepo mysql.OrderRepository) *CreateOrderService {
	return &CreateOrderService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run handles order creation logic
func (s *CreateOrderService) Run(req *payment.CreateOrderRequest) (*payment.OrderResponse, error) {
	// Validate request parameters
	if req.UserId == "" || req.Amount <= 0 || req.PaymentMethod == "" {
		return nil, errors.New("invalid request parameters")
	}

	// Create order object
	order := &models.PaymentOrder{
		UserID:        req.UserId,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Insert order into database
	createdOrder, err := s.orderRepo.CreateOrder(s.ctx, order)
	if err != nil {
		return nil, err
	}

	// Construct response
	resp := &payment.OrderResponse{
		OrderId:       createdOrder.OrderID,
		UserId:        createdOrder.UserID,
		Amount:        createdOrder.Amount,
		PaymentMethod: createdOrder.PaymentMethod,
		Status:        createdOrder.Status,
		CreatedAt:     createdOrder.CreatedAt.Unix(),
	}

	return resp, nil
}
