package service

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/infra/rpc"
	product "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/product"
	rpcProduct "github.com/MelodyDeep/TikTok-E-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
)

type DeleteProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteProductService(Context context.Context, RequestContext *app.RequestContext) *DeleteProductService {
	return &DeleteProductService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteProductService) Run(req *product.DeleteProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	r, err := rpc.ProductClient.DeleteProduct(h.Context, &rpcProduct.DeleteProductReq{Id: req.ProductId, StoreId: req.StoreId})
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"success": r.Success,
	}, nil
}
