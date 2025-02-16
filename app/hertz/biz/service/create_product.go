package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/product"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type CreateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateProductService(Context context.Context, RequestContext *app.RequestContext) *CreateProductService {
	return &CreateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateProductService) Run(req *product.CreateProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res, err := rpc.ProductClient.CreateProduct(h.Context, &rpcProduct.CreateProductReq{Product: &rpcProduct.Product{
		StoreId: req.StoreId,
		Name: req.Name,
		Description: req.Description,
		Picture: req.Picture,
		Price: req.Price,
		Stock: req.Stock,
		Categories: req.Categories,
	}})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"id": res.Id,
	}, nil
}
