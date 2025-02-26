package payment_gateway

import "fmt"

// AliPay 实现了 PaymentStrategy 接口，表示支付宝支付方式
type AliPay struct{}

// Pay 支付宝支付的实现
func (a *AliPay) Pay(paymentReq PaymentRequest) (string, error) {
	// 调用支付宝支付API
	return fmt.Sprintf("AliPay Payment for Order %s successful", paymentReq.OrderID), nil
}
