package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateOrderStatusService struct {
	ctx context.Context
} // NewUpdateOrderStatusService new UpdateOrderStatusService
func NewUpdateOrderStatusService(ctx context.Context) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderStatusService) Run(req *order_service.UpdateOrderStatusRequest) (resp *order_service.BaseResponse, err error) {
	// Finish your business logic.

	orderDAO := dao.GetOrderDAO()
	order, err := orderDAO.FindOne(req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	order.Status = req.Status
	err = orderDAO.Update(order)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return &order_service.BaseResponse{
		Code: 200,
		Msg:  "ok",
	}, nil
}
