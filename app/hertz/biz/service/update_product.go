package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/product"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type UpdateProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateProductService(Context context.Context, RequestContext *app.RequestContext) *UpdateProductService {
	return &UpdateProductService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateProductService) Run(req *product.UpdateProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res, err := rpc.ProductClient.UpdateProduct(h.Context, &rpcProduct.UpdateProductReq{
		Product: &rpcProduct.Product{
			Id:          req.Id,
			StoreId:     req.StoreId,
			Name:        req.Name,
			Description: req.Description,
			Picture:     req.Picture,
			Price:       req.Price,
			Stock:       req.Stock,
			Categories:  req.Categories,
		},
	})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"success": res.Success,
	}, nil
}
