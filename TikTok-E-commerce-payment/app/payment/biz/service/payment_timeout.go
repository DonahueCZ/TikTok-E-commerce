package service

import (
	"context"
	"time"
)

// HandlePaymentTimeout 处理支付超时
func (s *orderService) HandlePaymentTimeout(ctx context.Context, orderID string) error {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}
	if order == nil {
		return nil // 订单不存在，直接返回
	}

	// 2. 检查订单是否超时（假设超时时间是 30 分钟）
	if time.Since(order.CreatedAt) < 30*time.Minute {
		return nil // 未超时，不做处理
	}

	// 3. 更新订单状态为 "timeout"
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, "timeout")
}
