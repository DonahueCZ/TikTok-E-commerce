package rpc

import (
	"TikTok-E-commerce-payment/kitex_gen/payment_proto_idl/idl/paymentservice/paymentservice"
	"context"
	"testing"
)

func TestProcessPayment(t *testing.T) {
	s := &PaymentServiceImpl{}

	req := &paymentservice.PaymentRequest{
		OrderId: "test-order",
	}

	resp, err := s.ProcessPayment(context.Background(), req)
	if err != nil {
		t.Fatalf("ProcessPayment 调用失败: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("期望 Status=success，得到: %s", resp.Status)
	}

}
