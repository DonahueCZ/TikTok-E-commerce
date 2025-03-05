package service

import (
	"context"
	"errors"
	"time"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/models"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type CreateOrderService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewCreateOrderService 创建 CreateOrderService
func NewCreateOrderService(ctx context.Context, orderRepo mysql.OrderRepository) *CreateOrderService {
	return &CreateOrderService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 处理订单创建逻辑
func (s *CreateOrderService) Run(req *payment.CreateOrderRequest) (*payment.OrderResponse, error) {
	// 验证请求参数
	if req.OrderId == "" || req.Amount <= 0 || req.PaymentMethod == "" {
		return nil, errors.New("请求参数无效")
	}

	// 创建订单对象
	order := &models.PaymentOrder{
		OrderID:       req.OrderId,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Status:        "pending",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// 插入订单
	createdOrderID, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	// 返回响应
	return &payment.OrderResponse{
		OrderId: createdOrderID,
		Amount:  order.Amount,
		Status:  order.Status,
	}, nil
}
