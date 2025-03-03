package service

import (
	"context"
	payment "github.com/MelodyDeep/TikTok-E-commerce/app/payment/kitex_gen/rpc/payment"
	"testing"
)

func TestCancelPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelPaymentService(ctx)
	// init req and assert value

	req := &payment.CancelRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
