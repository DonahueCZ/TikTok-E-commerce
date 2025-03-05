package service

import (
	"context"
	"testing"
	payment "github.com/MelodyDeep/TikTok-E-commerce/app/payment/kitex_gen/rpc/payment"
)

func TestDeleteOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteOrderService(ctx)
	// init req and assert value

	req := &payment.DeleteOrderRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
