package service

// 负责用户主动取消支付

import "context"

func (s *orderService) CancelPayment(ctx context.Context, orderID string) error {
	// 1. 查询订单
	order, err := s.orderRepo.GetOrderByID(ctx, orderID)
	if err != nil || order == nil {
		return err
	}

	// 2. 检查订单状态
	if order.Status == "已支付" {
		return errors.New("订单已支付，无法取消")
	}

	// 3. 更新订单状态
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, "已取消")
}
