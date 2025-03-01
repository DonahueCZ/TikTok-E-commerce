package handler

import (
	"context"
	"testing"

	"TikTok-E-commerce-payment/app/payment/biz/service"
	"TikTok-E-commerce-payment/kitex_gen/paymentservice/paymentservice"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Service
type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) ProcessPayment(ctx context.Context, payment *service.Payment) (*service.Payment, error) {
	args := m.Called(ctx, payment)
	return args.Get(0).(*service.Payment), args.Error(1)
}

// ✅ 测试 `ProcessPayment`
func TestProcessPayment(t *testing.T) {
	mockService := new(MockOrderService)
	handler := NewOrderHandler(mockService)

	req := &paymentservice.ProcessPaymentRequest{
		OrderId:       "test_order_123",
		Amount:        100.5,
		PaymentMethod: "Alipay",
	}

	mockPayment := &service.Payment{
		OrderID: req.OrderId,
		Amount:  req.Amount,
		Status:  "paid",
	}

	mockService.On("ProcessPayment", mock.Anything, mock.Anything).Return(mockPayment, nil)

	resp, err := handler.ProcessPayment(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "支付成功", resp.Message)
	mockService.AssertExpectations(t)
}
