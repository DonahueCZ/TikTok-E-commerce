package main

import (
	"context"
	"order/rpc/order/db/dao"
	"order/rpc/order/db/ordermd"
	order_service "order/rpc/order/kitex_gen/demo/order_service"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.BaseResponse, err error) {
	// TODO: Your code here...
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

// QueryOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrderById(ctx context.Context, req *order_service.QueryOrderByIdRequest) (resp *order_service.QueryOrderResponse, err error) {
	// TODO: Your code here...
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

// QueryOrdersByUserId implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) QueryOrdersByUserId(ctx context.Context, req *order_service.QueryOrdersByUserIdRequest) (resp *order_service.QueryOrdersResponse, err error) {
	// TODO: Your code here...
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

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *order_service.BaseResponse, err error) {
	// TODO: Your code here...
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

// UpdateOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *order_service.UpdateOrderStatusRequest) (resp *order_service.BaseResponse, err error) {
	// TODO: Your code here...
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

// UpdateOrderAddresseeInfo implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderAddresseeInfo(ctx context.Context, req *order_service.UpdateOrderAddresseeInfoRequest) (resp *order_service.BaseResponse, err error) {
	// TODO: Your code here...
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
