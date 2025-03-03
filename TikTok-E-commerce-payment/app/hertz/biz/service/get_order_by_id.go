package service

import (
	"context"
	"fmt"

	rpcPayment "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetOrderByIDService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetOrderByIDService(Context context.Context, RequestContext *app.RequestContext) *GetOrderByIDService {
	return &GetOrderByIDService{RequestContext: RequestContext, Context: Context}
}

func (h *GetOrderByIDService) Run(req *rpcPayment.GetOrderByIDRequest) (*rpcPayment.OrderResponse, error) {
	rpcClient := rpcPayment.NewPaymentServiceClient("etcd://127.0.0.1:2379/payment") // etcd 地址
	rpcResp, err := rpcClient.GetOrderByID(h.Context, req)
	if err != nil {
		fmt.Println("查询订单失败:", err)
		return nil, err
	}

	resp := &rpcPayment.OrderResponse{
		OrderId: rpcResp.OrderId,
		Status:  rpcResp.Status,
	}
	return resp, nil
}
