package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/product"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	list, err := rpc.ProductClient.SearchProducts(h.Context, &rpcProduct.SearchProductsReq{Query: req.Query})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"products": list.Results,
		"query":    req.Query,
	}, nil
}
