package service

import (
	"context"
	order "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	rpcclient "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/order_rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateOrderStatusService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderStatusService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderStatusService) Run(req *order.OrderStatusRequest) (resp *order.BaseResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcclient.InitOrderRpcClient()

	_, err = orderrpc.UpdateOrderStatus(h.Context, &order_service.UpdateOrderStatusRequest{
		OrderId: req.OrderId,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	resp = new(order.BaseResponse)
	resp.Code = 200
	resp.Message = "ok"
	return resp, nil
}
