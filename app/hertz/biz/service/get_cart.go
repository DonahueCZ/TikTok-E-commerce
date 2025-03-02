package service

import (
	"context"
	"strconv"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	// common "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/common"
	cart "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/cart"
	rpcCart "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/cart"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *cart.GetCartReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

    var rpcResp *rpcCart.GetCartResp
	rpcResp, err = rpc.CartClient.GetCart(h.Context, &rpcCart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
    
	cartResp := rpcResp
	var items []map[string]string
	var total float64
	for _,item := range cartResp.Items{
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcProduct.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
	}
	p := productResp.Product
	items = append(items, map[string]string{
		"name":        p.Name,
		"Description": p.Description,
		"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
		"Picture":     p.Picture,
		"Qty":         strconv.Itoa(int(item.Quantity)),
	})
	total += float64(p.Price) * float64(item.Quantity)
}
return utils.H{
	"title": "Get Cart",
	"items":  items,
	"total":  total,
},nil
}
