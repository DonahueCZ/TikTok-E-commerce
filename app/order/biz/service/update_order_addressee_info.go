package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
)

type UpdateOrderAddresseeInfoService struct {
	ctx context.Context
} // NewUpdateOrderAddresseeInfoService new UpdateOrderAddresseeInfoService
func NewUpdateOrderAddresseeInfoService(ctx context.Context) *UpdateOrderAddresseeInfoService {
	return &UpdateOrderAddresseeInfoService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderAddresseeInfoService) Run(req *order_service.UpdateOrderAddresseeInfoRequest) (resp *order_service.BaseResponse, err error) {
	// Finish your business logic.
	
	orderDAO := dao.GetOrderDAO()
	order, err := orderDAO.FindOne(req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	order.AddresseeInfo, err = ordermd.AddresseeInfo2Str(req.AddresseeInfo)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	err = orderDAO.Update(order)
	if err != nil {
		klog.Error(err)
	}
	return &order_service.BaseResponse{
		Code: 200,
		Msg:  "ok",
	}, nil
}
