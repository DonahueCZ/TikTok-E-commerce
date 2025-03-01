package service

import (
	"context"
	"errors"
)

// CancelPayment 取消支付
func (s *orderService) CancelPayment(ctx context.Context, orderID string) error {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return errors.New("订单不存在")
	}

	// 2. 订单已经支付，无法取消
	if order.Status == "paid" {
		return errors.New("订单已支付，无法取消")
	}

	// 3. 更新订单状态为 "canceled"
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, "canceled")
}
