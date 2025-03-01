package service

import (
	"context"
)

// UpdateOrderStatus 更新订单状态
func (s *orderService) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}
