package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type GetOrderByIDService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewGetOrderByIDService 创建 GetOrderByIDService
func NewGetOrderByIDService(ctx context.Context, orderRepo mysql.OrderRepository) *GetOrderByIDService {
	return &GetOrderByIDService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 查询订单
func (s *GetOrderByIDService) Run(req *payment.OrderRequest) (*payment.OrderResponse, error) {
	// 查询订单
	order, err := s.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 返回响应
	return &payment.OrderResponse{
		OrderId: order.OrderID,
		Amount:  order.Amount,
		Status:  order.Status,
	}, nil // ⚠️ **删除 `PaymentMethod`（OrderResponse 里没有这个字段）**
}
