package service

import (
	"context"
	order "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	rpcclient "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/order_rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateOrderAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateOrderAddressService(Context context.Context, RequestContext *app.RequestContext) *UpdateOrderAddressService {
	return &UpdateOrderAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateOrderAddressService) Run(req *order.OrderAddressRequest) (resp *order.BaseResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcclient.InitOrderRpcClient()

	_, err = orderrpc.UpdateOrderAddresseeInfo(h.Context, &order_service.UpdateOrderAddresseeInfoRequest{
		OrderId: req.OrderId,
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
