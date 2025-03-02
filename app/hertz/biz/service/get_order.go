package service

import (
	"context"
	order "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	rpcclient "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/order_rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetOrderService(Context context.Context, RequestContext *app.RequestContext) *GetOrderService {
	return &GetOrderService{RequestContext: RequestContext, Context: Context}
}

func (h *GetOrderService) Run(req *order.OrderRequest) (resp *order.OrderResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	rpcclient.InitOrderRpcClient()

	res, err := orderrpc.QueryOrderById(h.Context, &order_service.QueryOrderByIdRequest{
		OrderId: req.OrderId,
	})

	if err != nil {

		return nil, err
	}

	resp = &order.OrderResponse{
		Order: &order.Order{
			ID:         res.OrderResponse.OrderId,
			UserId:     res.OrderResponse.UserId,
			GoodsId:    res.OrderResponse.GoodsId,
			GoodsCount: res.OrderResponse.GoodsCount,
			Cost:       res.OrderResponse.Cost,
			Status:     res.OrderResponse.Status,
			Address: &order.Address{
				Name:    res.OrderResponse.AddresseeInfo.Name,
				Phone:   res.OrderResponse.AddresseeInfo.Phone,
				Address: res.OrderResponse.AddresseeInfo.Address,
			},
			CreateTime: res.OrderResponse.CreateTime,
		},
	}
	return resp, nil
}
