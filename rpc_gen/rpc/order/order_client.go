package order

import (
	"context"
	order_service "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"

	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	CreateOrder(ctx context.Context, Req *order_service.CreateOrderRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error)
	QueryOrderById(ctx context.Context, Req *order_service.QueryOrderByIdRequest, callOptions ...callopt.Option) (r *order_service.QueryOrderResponse, err error)
	QueryOrdersByUserId(ctx context.Context, Req *order_service.QueryOrdersByUserIdRequest, callOptions ...callopt.Option) (r *order_service.QueryOrdersResponse, err error)
	UpdateOrder(ctx context.Context, Req *order_service.UpdateOrderRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error)
	UpdateOrderStatus(ctx context.Context, Req *order_service.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error)
	UpdateOrderAddresseeInfo(ctx context.Context, Req *order_service.UpdateOrderAddresseeInfoRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateOrder(ctx context.Context, Req *order_service.CreateOrderRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error) {
	return c.kitexClient.CreateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) QueryOrderById(ctx context.Context, Req *order_service.QueryOrderByIdRequest, callOptions ...callopt.Option) (r *order_service.QueryOrderResponse, err error) {
	return c.kitexClient.QueryOrderById(ctx, Req, callOptions...)
}

func (c *clientImpl) QueryOrdersByUserId(ctx context.Context, Req *order_service.QueryOrdersByUserIdRequest, callOptions ...callopt.Option) (r *order_service.QueryOrdersResponse, err error) {
	return c.kitexClient.QueryOrdersByUserId(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrder(ctx context.Context, Req *order_service.UpdateOrderRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error) {
	return c.kitexClient.UpdateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrderStatus(ctx context.Context, Req *order_service.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error) {
	return c.kitexClient.UpdateOrderStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrderAddresseeInfo(ctx context.Context, Req *order_service.UpdateOrderAddresseeInfoRequest, callOptions ...callopt.Option) (r *order_service.BaseResponse, err error) {
	return c.kitexClient.UpdateOrderAddresseeInfo(ctx, Req, callOptions...)
}
