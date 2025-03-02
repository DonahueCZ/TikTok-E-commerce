package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"testing"
)

func TestCreateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateOrderService(ctx)
	// init req and assert value

	req := &order_service.CreateOrderRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
