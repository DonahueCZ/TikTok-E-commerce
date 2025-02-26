package payment_gateway

import "fmt"

// BankCard 实现了 PaymentStrategy 接口，表示银行卡支付方式
type BankCard struct{}

// Pay 银行卡支付的实现
func (b *BankCard) Pay(paymentReq PaymentRequest) (string, error) {
	// 调用银行卡支付API
	return fmt.Sprintf("BankCard Payment for Order %s successful", paymentReq.OrderID), nil
}
