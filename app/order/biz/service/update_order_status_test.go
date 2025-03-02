package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	"testing"
)

func TestUpdateOrderStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateOrderStatusService(ctx)
	// init req and assert value

	req := &order_service.UpdateOrderStatusRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
