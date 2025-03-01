package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
)

type QueryOrderByIdService struct {
	ctx context.Context
} // NewQueryOrderByIdService new QueryOrderByIdService
func NewQueryOrderByIdService(ctx context.Context) *QueryOrderByIdService {
	return &QueryOrderByIdService{ctx: ctx}
}

// Run create note info
func (s *QueryOrderByIdService) Run(req *order_service.QueryOrderByIdRequest) (resp *order_service.QueryOrderResponse, err error) {
	// Finish your business logic.

	orderDAO := dao.GetOrderDAO()
	order, err := orderDAO.FindOne(req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	addresseeInfo, err := ordermd.Str2AddresseeInfo(order.AddresseeInfo)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	return &order_service.QueryOrderResponse{
		OrderResponse: &order_service.OrderResponse{
			OrderId:       order.Id,
			UserId:        order.UserId,
			GoodsId:       order.GoodsId,
			Status:        order.Status,
			CreateTime:    order.CreateTime,
			GoodsCount:    order.GoodsCount,
			Cost:          order.Cost,
			AddresseeInfo: addresseeInfo,
		},
	}, nil
}
