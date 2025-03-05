package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type DeleteOrderService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewDeleteOrderService 创建 DeleteOrderService
func NewDeleteOrderService(ctx context.Context, orderRepo mysql.OrderRepository) *DeleteOrderService {
	return &DeleteOrderService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 处理删除订单逻辑
func (s *DeleteOrderService) Run(req *payment.DeleteOrderRequest) (*payment.PaymentResponse, error) {
	// 1. 先查询订单是否存在
	order, err := s.orderRepo.GetOrderByID(req.OrderId)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在，无法删除")
	}

	// 2. 订单已支付的情况下可能不允许删除（视业务需求而定）
	if order.Status == "paid" {
		return nil, errors.New("订单已支付，无法删除")
	}

	// 3. 调用 repository 层删除订单
	err = s.orderRepo.DeleteOrder(req.OrderId) // ⚠️ 这里修正 `mysql.OrderRepository` 变量调用
	if err != nil {
		return nil, err
	}

	// 4. 返回成功响应
	return &payment.PaymentResponse{
		Status:  "success",
		Message: "订单删除成功",
	}, nil
}
