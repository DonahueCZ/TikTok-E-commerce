package service

import (
	"context"
	"errors"

	"github.com/MelodyDeep/TikTok-E-commerce/app/payment/biz/dal/mysql"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
)

type UpdateOrderStatusService struct {
	ctx       context.Context
	orderRepo mysql.OrderRepository
}

// NewUpdateOrderStatusService 创建 UpdateOrderStatusService
func NewUpdateOrderStatusService(ctx context.Context, orderRepo mysql.OrderRepository) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{
		ctx:       ctx,
		orderRepo: orderRepo,
	}
}

// Run 更新订单状态
func (s *UpdateOrderStatusService) Run(req *payment.UpdateStatusRequest) (*payment.PaymentResponse, error) {
	// 1. 查询订单是否存在
	order, err := s.orderRepo.GetOrderByID(req.OrderId) // ⚠️ **修正 `mysql.OrderRepo`，改为 `s.orderRepo`**
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 订单状态检查，避免非法状态转换
	if order.Status == req.NewStatus { // ⚠️ **修正 `req.Status`，改为 `req.NewStatus`**
		return &payment.PaymentResponse{
			Status:  "failed", // ⚠️ **修正 `Success`，改为 `Status`**
			Message: "订单状态未变化",
		}, nil
	}

	// 3. 更新订单状态
	err = s.orderRepo.UpdateOrderStatus(req.OrderId, req.NewStatus) // ⚠️ **修正 `req.Status`，改为 `req.NewStatus`**
	if err != nil {
		return nil, err
	}

	// 4. 返回更新成功的响应
	return &payment.PaymentResponse{
		Status:  "success", // ⚠️ **修正 `Success`，改为 `Status`**
		Message: "订单状态更新成功",
	}, nil
}
