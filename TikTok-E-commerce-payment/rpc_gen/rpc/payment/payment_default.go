package payment

import (
	"context"
	payment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/rpc/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ProcessPayment(ctx context.Context, req *payment.PaymentRequest, callOptions ...callopt.Option) (resp *payment.PaymentResponse, err error) {
	resp, err = defaultClient.ProcessPayment(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ProcessPayment call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelPayment(ctx context.Context, req *payment.CancelRequest, callOptions ...callopt.Option) (resp *payment.PaymentResponse, err error) {
	resp, err = defaultClient.CancelPayment(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelPayment call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func HandlePaymentTimeout(ctx context.Context, req *payment.PaymentTimeoutRequest, callOptions ...callopt.Option) (resp *payment.PaymentResponse, err error) {
	resp, err = defaultClient.HandlePaymentTimeout(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "HandlePaymentTimeout call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrderByID(ctx context.Context, req *payment.OrderRequest, callOptions ...callopt.Option) (resp *payment.OrderResponse, err error) {
	resp, err = defaultClient.GetOrderByID(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrderByID call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrderStatus(ctx context.Context, req *payment.UpdateStatusRequest, callOptions ...callopt.Option) (resp *payment.PaymentResponse, err error) {
	resp, err = defaultClient.UpdateOrderStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrderStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteOrder(ctx context.Context, req *payment.DeleteOrderRequest, callOptions ...callopt.Option) (resp *payment.PaymentResponse, err error) {
	resp, err = defaultClient.DeleteOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateOrder(ctx context.Context, req *payment.CreateOrderRequest, callOptions ...callopt.Option) (resp *payment.OrderResponse, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
