package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order_service.UpdateOrderRequest) (resp *order_service.BaseResponse, err error) {
	// Finish your business logic.
	orderDAO := dao.GetOrderDAO()
	order, err := orderDAO.FindOne(req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	order.Status = req.Status
	order.GoodsCount = req.GoodsCount
	order.Cost = req.Cost
	order.AddresseeInfo, err = ordermd.AddresseeInfo2Str(req.AddresseeInfo)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
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
