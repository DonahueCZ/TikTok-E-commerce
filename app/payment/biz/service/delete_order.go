package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce_payment/rpc_gen/kitex_gen/rpc/payment"
)

type DeleteOrderService struct {
	ctx context.Context
}

// NewDeleteOrderService 创建 DeleteOrderService
func NewDeleteOrderService(ctx context.Context) *DeleteOrderService {
	return &DeleteOrderService{ctx: ctx}
}

// Run 处理删除订单逻辑
func (s *DeleteOrderService) Run(req *payment.DeleteOrderRequest) (resp *payment.PaymentResponse, err error) {
	// 1. 先查询订单是否存在
	order, err := mysql.OrderRepo.GetOrderByID(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	// 2. 如果订单不存在，返回错误
	if order == nil {
		return nil, errors.New("订单不存在，无法删除")
	}

	// 3. 订单已支付的情况下可能不允许删除（视业务需求而定）
	if order.Status == "paid" {
		return nil, errors.New("订单已支付，无法删除")
	}

	// 4. 调用 repository 层删除订单
	err = mysql.OrderRepo.DeleteOrder(s.ctx, req.OrderId)
	if err != nil {
		return nil, err
	}

	// 5. 返回成功响应
	resp = &payment.PaymentResponse{
		Success: true,
		Message: "订单删除成功",
	}
	return resp, nil
}
