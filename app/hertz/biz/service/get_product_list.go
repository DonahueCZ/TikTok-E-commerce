package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/product"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

const pageSize = 10

type GetProductListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductListService(Context context.Context, RequestContext *app.RequestContext) *GetProductListService {
	return &GetProductListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductListService) Run(req *product.GetProductListReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	if req.Page == 0 {
		req.Page = 1
	}
	list, err := rpc.ProductClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{Page: req.Page, PageSize: pageSize, CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"products": list.Products,
		"page": req.Page,
		"category": req.Category,
	}, nil
}
