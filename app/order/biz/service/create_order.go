package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

type CreateOrderService struct {
	ctx context.Context
} // NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context) *CreateOrderService {
	return &CreateOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateOrderService) Run(req *order_service.CreateOrderRequest) (resp *order_service.BaseResponse, err error) {
	// Finish your business logic.

	orderDAO := dao.GetOrderDAO()
	addrStr, err := ordermd.AddresseeInfo2Str(req.AddresseeInfo)
	if err != nil {
		klog.Error(err)
		return &order_service.BaseResponse{
			Code: 1001,
			Msg:  "AddresseeInfo2Str fail",
		}, err
	}
	data := &ordermd.Order{
		UserId:        req.UserId,
		GoodsId:       req.GoodsId,
		Status:        ordermd.IsNotPaid,
		CreateTime:    time.Now().Unix(),
		GoodsCount:    req.GoodsCount,
		Cost:          req.Cost,
		AddresseeInfo: addrStr,
	}
	err = orderDAO.Insert(data)
	if err != nil {
		klog.Error(err)
		return &order_service.BaseResponse{
			Code: 10010,
			Msg:  err.Error(),
		}, err
	}
	return &order_service.BaseResponse{
		Code: 200,
		Msg:  "ok",
	}, nil
}
