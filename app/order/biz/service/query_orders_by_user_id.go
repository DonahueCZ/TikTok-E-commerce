package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/dao"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/dal/ordermd"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/pkg/klog"
)

type QueryOrdersByUserIdService struct {
	ctx context.Context
} // NewQueryOrdersByUserIdService new QueryOrdersByUserIdService
func NewQueryOrdersByUserIdService(ctx context.Context) *QueryOrdersByUserIdService {
	return &QueryOrdersByUserIdService{ctx: ctx}
}

// Run create note info
func (s *QueryOrdersByUserIdService) Run(req *order_service.QueryOrdersByUserIdRequest) (resp *order_service.QueryOrdersResponse, err error) {
	// Finish your business logic.

	orderDAO := dao.GetOrderDAO()
	orders, err := orderDAO.FindByUserId(req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	orderResponses := make([]*order_service.OrderResponse, 0, len(orders))
	for _, order := range orders {
		addresseeInfo, err := ordermd.Str2AddresseeInfo(order.AddresseeInfo)
		if err != nil {
			return nil, err
		}
		orderResponses = append(orderResponses,
			&order_service.OrderResponse{
				OrderId:       order.Id,
				UserId:        order.UserId,
				GoodsId:       order.GoodsId,
				Status:        order.Status,
				CreateTime:    order.CreateTime,
				GoodsCount:    order.GoodsCount,
				Cost:          order.Cost,
				AddresseeInfo: addresseeInfo,
			})
	}
	return &order_service.QueryOrdersResponse{
		OrderResponses: orderResponses,
	}, nil
}
