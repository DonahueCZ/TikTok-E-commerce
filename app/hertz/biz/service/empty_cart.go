package service

import (
	"context"

	cart "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/cart"
	common "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/common"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
