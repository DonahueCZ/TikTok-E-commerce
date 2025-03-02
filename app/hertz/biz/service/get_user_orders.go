package service

import (
	"context"
	order "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	rpcclient "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/rpc_client/order_rpc"
	"github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/order_service"
	orderrpc "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/rpc/order"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserOrdersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserOrdersService(Context context.Context, RequestContext *app.RequestContext) *GetUserOrdersService {
	return &GetUserOrdersService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserOrdersService) Run(req *order.UserOrderRequest) (resp *order.OrderListResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	rpcclient.InitOrderRpcClient()

	res, err := orderrpc.QueryOrdersByUserId(h.Context, &order_service.QueryOrdersByUserIdRequest{
		UserId:   req.UserId,
		PageSize: req.Size,
		PageNum:  req.Page,
	})

	if err != nil {
		return nil, err
	}
	orders := make([]*order.Order, 0)
	for _, o := range res.OrderResponses {
		orders = append(orders, &order.Order{
			ID:         o.OrderId,
			UserId:     o.UserId,
			GoodsId:    o.GoodsId,
			GoodsCount: o.GoodsCount,
			Cost:       o.Cost,
			Status:     o.Status,
			Address: &order.Address{
				Name:    o.AddresseeInfo.Name,
				Phone:   o.AddresseeInfo.Phone,
				Address: o.AddresseeInfo.Address,
			},
			CreateTime: o.CreateTime,
		})
	}
	resp = &order.OrderListResponse{
		Orders: orders,
	}

	return resp, nil
}
