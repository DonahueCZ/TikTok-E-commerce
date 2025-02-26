package payment_gateway

import "fmt"

// WeChatPay 实现了 PaymentStrategy 接口，表示微信支付方式
type WeChatPay struct{}

// Pay 微信支付的实现
func (w *WeChatPay) Pay(paymentReq PaymentRequest) (string, error) {
	// 调用微信支付API
	return fmt.Sprintf("WeChat Payment for Order %s successful", paymentReq.OrderID), nil
}
