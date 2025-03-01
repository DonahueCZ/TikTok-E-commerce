package handler

import (
	"context"

	"TikTok-E-commerce-payment/app/payment/biz/service"
	"TikTok-E-commerce-payment/kitex_gen/paymentservice/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
)

// ✅ OrderHandler 依赖 OrderService
type OrderHandler struct {
	service service.OrderService
}

// ✅ 修改构造函数，接收 `OrderService`
func NewOrderHandler(svc service.OrderService) *OrderHandler {
	return &OrderHandler{service: svc}
}

// ✅ 处理支付请求
func (h *OrderHandler) ProcessPayment(ctx context.Context, req *paymentservice.ProcessPaymentRequest) (*paymentservice.ProcessPaymentResponse, error) {
	klog.Infof("处理支付请求: OrderID=%s, Amount=%.2f, PaymentMethod=%s",
		req.OrderId, req.Amount, req.PaymentMethod)

	payment := &service.Payment{
		OrderID: req.OrderId,
		Amount:  req.Amount,
		Status:  "pending",
	}

	paidPayment, err := h.service.ProcessPayment(ctx, payment)
	if err != nil {
		return &paymentservice.ProcessPaymentResponse{
			Status:  "failed",
			Message: err.Error(),
		}, err
	}

	return &paymentservice.ProcessPaymentResponse{
		Status:  "success",
		Message: "支付成功",
	}, nil
}
