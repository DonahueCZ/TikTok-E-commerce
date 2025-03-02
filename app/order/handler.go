package main

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/order/biz/service"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.BaseResponse, err error) {
	resp, err = service.NewCreateOrderService(ctx).Run(req)

	return resp, err
}

// QueryOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrderById(ctx context.Context, req *order_service.QueryOrderByIdRequest) (resp *order_service.QueryOrderResponse, err error) {
	resp, err = service.NewQueryOrderByIdService(ctx).Run(req)

	return resp, err
}

// QueryOrdersByUserId implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrdersByUserId(ctx context.Context, req *order_service.QueryOrdersByUserIdRequest) (resp *order_service.QueryOrdersResponse, err error) {
	resp, err = service.NewQueryOrdersByUserIdService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *order_service.BaseResponse, err error) {
	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *order_service.UpdateOrderStatusRequest) (resp *order_service.BaseResponse, err error) {
	resp, err = service.NewUpdateOrderStatusService(ctx).Run(req)

	return resp, err
}

// UpdateOrderAddresseeInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderAddresseeInfo(ctx context.Context, req *order_service.UpdateOrderAddresseeInfoRequest) (resp *order_service.BaseResponse, err error) {
	resp, err = service.NewUpdateOrderAddresseeInfoService(ctx).Run(req)

	return resp, err
}
