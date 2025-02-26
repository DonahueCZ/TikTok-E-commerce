package service

// 负责定时取消支付
import (
	"context"
	"time"
)

func (s *orderService) HandlePaymentTimeout(ctx context.Context, orderID string) error {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil || order == nil {
		return err
	}

	// 2. 检查订单是否超时（假设超时时间是30分钟）
	if time.Since(order.CreatedAt) < 30*time.Minute {
		return nil // 未超时，不做处理
	}

	// 3. 更新订单状态
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, "超时未支付")
}
