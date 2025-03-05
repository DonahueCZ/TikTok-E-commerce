package payment

import (
	"context"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"

	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	ProcessPayment(ctx context.Context, Req *payment.PaymentRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error)
	CancelPayment(ctx context.Context, Req *payment.CancelRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error)
	HandlePaymentTimeout(ctx context.Context, Req *payment.PaymentTimeoutRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error)
	GetOrderByID(ctx context.Context, Req *payment.OrderRequest, callOptions ...callopt.Option) (r *payment.OrderResponse, err error)
	UpdateOrderStatus(ctx context.Context, Req *payment.UpdateStatusRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error)
	DeleteOrder(ctx context.Context, Req *payment.DeleteOrderRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error)
	CreateOrder(ctx context.Context, Req *payment.CreateOrderRequest, callOptions ...callopt.Option) (r *payment.OrderResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ProcessPayment(ctx context.Context, Req *payment.PaymentRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error) {
	return c.kitexClient.ProcessPayment(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelPayment(ctx context.Context, Req *payment.CancelRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error) {
	return c.kitexClient.CancelPayment(ctx, Req, callOptions...)
}

func (c *clientImpl) HandlePaymentTimeout(ctx context.Context, Req *payment.PaymentTimeoutRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error) {
	return c.kitexClient.HandlePaymentTimeout(ctx, Req, callOptions...)
}

func (c *clientImpl) GetOrderByID(ctx context.Context, Req *payment.OrderRequest, callOptions ...callopt.Option) (r *payment.OrderResponse, err error) {
	return c.kitexClient.GetOrderByID(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrderStatus(ctx context.Context, Req *payment.UpdateStatusRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error) {
	return c.kitexClient.UpdateOrderStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteOrder(ctx context.Context, Req *payment.DeleteOrderRequest, callOptions ...callopt.Option) (r *payment.PaymentResponse, err error) {
	return c.kitexClient.DeleteOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateOrder(ctx context.Context, Req *payment.CreateOrderRequest, callOptions ...callopt.Option) (r *payment.OrderResponse, err error) {
	return c.kitexClient.CreateOrder(ctx, Req, callOptions...)
}
