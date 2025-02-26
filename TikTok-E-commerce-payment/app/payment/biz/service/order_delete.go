package service

import (
	"Errors"
	"context"
)

func (s *orderService) DeleteOrder(ctx context.Context, orderID string) error {
	// 1. 先查询订单是否存在
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	// 2. 订单存在，则调用 repository 层删除
	if order != nil {
		return s.orderRepo.DeleteOrder(ctx, orderID)
	}

	// 3. 订单不存在，返回错误
	return errors.New("订单不存在，无法删除")
}
