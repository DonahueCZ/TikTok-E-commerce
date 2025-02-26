package payment_gateway

import (
	_ "errors"
	_ "fmt"
)

// PaymentProcessor 处理支付请求
type PaymentProcessor struct {
	strategy PaymentStrategy // 支付策略接口
}

// NewPaymentProcessor 创建一个新的支付处理器
func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{
		strategy: strategy,
	}
}

// ProcessPayment 处理支付请求
func (p *PaymentProcessor) ProcessPayment(paymentReq PaymentRequest) (string, error) {
	// 调用相应的支付策略来处理支付
	return p.strategy.Pay(paymentReq)
}
