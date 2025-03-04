package payment

import (
	"context"

	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/service"
	"github.com/MelodyDeep/TikTok-E-commerce/app/hertz/biz/utils"
	payment "github.com/MelodyDeep/TikTok-E-commerce/app/hertz/hertz_gen/hertz/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CreateOrder .
// @router /payment/order/create [POST]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.CreateOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.OrderResponse{}
	resp, err = service.NewCreateOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetOrderByID .
// @router /payment/order/{order_id} [GET]
func GetOrderByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.GetOrderByIDRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.OrderResponse{}
	resp, err = service.NewGetOrderByIDService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateOrderStatus .
// @router /payment/update_status [POST]
func UpdateOrderStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.UpdateOrderStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.PaymentResponse{}
	resp, err = service.NewUpdateOrderStatusService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteOrder .
// @router /payment/delete [DELETE]
func DeleteOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.DeleteOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.PaymentResponse{}
	resp, err = service.NewDeleteOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ProcessPayment .
// @router /payment/process [POST]
func ProcessPayment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.ProcessPaymentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.PaymentResponse{}
	resp, err = service.NewProcessPaymentService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CancelPayment .
// @router /payment/cancel [POST]
func CancelPayment(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.CancelPaymentRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.PaymentResponse{}
	resp, err = service.NewCancelPaymentService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// HandlePaymentTimeout .
// @router /payment/timeout [POST]
func HandlePaymentTimeout(ctx context.Context, c *app.RequestContext) {
	var err error
	var req payment.HandlePaymentTimeoutRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &payment.PaymentResponse{}
	resp, err = service.NewHandlePaymentTimeoutService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
