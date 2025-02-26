package service

import (
	"context"
)

func (s *orderService) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}
