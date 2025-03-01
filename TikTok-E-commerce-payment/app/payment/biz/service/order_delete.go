package service

import (
	"context"
	"errors"
)

// DeleteOrder 删除订单
func (s *orderService) DeleteOrder(ctx context.Context, orderID string) error {
	// 先查询订单是否存在
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	// 订单不存在
	if order == nil {
		return errors.New("订单不存在，无法删除")
	}

	// 调用 repository 层删除订单
	return s.orderRepo.DeleteOrder(ctx, orderID)
}
