package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	cart "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/cart"
	rpcCart "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/cart"
	common "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddItemService(Context context.Context, RequestContext *app.RequestContext) *AddItemService {
	return &AddItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddItemService) Run(req *cart.AddItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_,err = rpc.CartClient.AddItem(h.Context, &rpcCart.AddItemReq{
		UserId:req.UserId,
		Item: &rpcCart.CartItem{
			ProductId: req.Item.ProductId,
			Quantity:  req.Item.Quantity,
},
})
if err != nil {
	return nil,err
}
	return
}
