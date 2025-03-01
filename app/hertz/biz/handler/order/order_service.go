package order

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/service"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/utils"
	order "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateOrder .
// @router /order [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.CreateOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.BaseResponse{}
	resp, err = service.NewCreateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetOrder .
// @router /order [GET]
func GetOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderResponse{}
	resp, err = service.NewGetOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetUserOrders .
// @router /user/orders [GET]
func GetUserOrders(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UserOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.OrderListResponse{}
	resp, err = service.NewGetUserOrdersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrder .
// @router /order [PUT]
func UpdateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.UpdateOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.BaseResponse{}
	resp, err = service.NewUpdateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrderStatus .
// @router /order/status [POST]
func UpdateOrderStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.BaseResponse{}
	resp, err = service.NewUpdateOrderStatusService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrderAddress .
// @router /order/address [POST]
func UpdateOrderAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req order.OrderAddressRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &order.BaseResponse{}
	resp, err = service.NewUpdateOrderAddressService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
