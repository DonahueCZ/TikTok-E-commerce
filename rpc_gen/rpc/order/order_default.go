package order

import (
	"context"
	order_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateOrder(ctx context.Context, req *order_service.CreateOrderRequest, callOptions ...callopt.Option) (resp *order_service.BaseResponse, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryOrderById(ctx context.Context, req *order_service.QueryOrderByIdRequest, callOptions ...callopt.Option) (resp *order_service.QueryOrderResponse, err error) {
	resp, err = defaultClient.QueryOrderById(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryOrderById call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func QueryOrdersByUserId(ctx context.Context, req *order_service.QueryOrdersByUserIdRequest, callOptions ...callopt.Option) (resp *order_service.QueryOrdersResponse, err error) {
	resp, err = defaultClient.QueryOrdersByUserId(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryOrdersByUserId call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrder(ctx context.Context, req *order_service.UpdateOrderRequest, callOptions ...callopt.Option) (resp *order_service.BaseResponse, err error) {
	resp, err = defaultClient.UpdateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrderStatus(ctx context.Context, req *order_service.UpdateOrderStatusRequest, callOptions ...callopt.Option) (resp *order_service.BaseResponse, err error) {
	resp, err = defaultClient.UpdateOrderStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrderStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrderAddresseeInfo(ctx context.Context, req *order_service.UpdateOrderAddresseeInfoRequest, callOptions ...callopt.Option) (resp *order_service.BaseResponse, err error) {
	resp, err = defaultClient.UpdateOrderAddresseeInfo(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrderAddresseeInfo call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
