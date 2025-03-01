package cart

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/service"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/utils"
	cart "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/cart"
	common "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddItem .
// @router /cart/items [POST]
func AddItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewAddItemService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetCart .
// @router /cart/:user_id [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.GetCartReq
	var resp map[string]any
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &cart.GetCartResp{}
	resp, err = service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// EmptyCart .
// @router /cart/:user_id [DELETE]
func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.EmptyCartReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &common.Empty{}
	resp, err = service.NewEmptyCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
