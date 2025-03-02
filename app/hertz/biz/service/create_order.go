package service

import (
	"context"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/order_rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateOrderService(Context context.Context, RequestContext *app.RequestContext) *CreateOrderService {
	return &CreateOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateOrderService) Run(req *order.CreateOrderRequest) (resp *order.BaseResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	orderrpcclient.InitOrderRpcClient()

	_, err = orderrpc.CreateOrder(h.Context, &order_service.CreateOrderRequest{
		UserId:     req.UserId,
		GoodsId:    req.GoodsId,
		GoodsCount: req.GoodsCount,
		Cost:       req.Cost,
		AddresseeInfo: &order_service.AddresseeInfo{
			Name:    req.Address.Name,
			Phone:   req.Address.Phone,
			Address: req.Address.Address,
		},
	})

	if err != nil {
		return nil, err
	}

	resp = new(order.BaseResponse)
	resp.Code = 200
	resp.Message = "ok"
	return resp, nil
}
