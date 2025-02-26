package payment_gateway

import (
	"errors"
)

const (
	// 支付方式常量
	WeChatPayType = "wechat"
	AliPayType    = "alipay"
	BankCardType  = "bankcard"
)

type PaymentRequest struct {
	OrderID     string
	Amount      float64
	PaymentType string
}

// PaymentStrategy 支付方式策略接口
type PaymentStrategy interface {
	Pay(paymentReq PaymentRequest) (string, error)
}

// PaymentFactory 用于获取支付策略
type PaymentFactory struct{}

// NewPaymentFactory 创建一个支付方式工厂
func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{}
}

// GetPaymentStrategy 根据支付方式选择策略
func (f *PaymentFactory) GetPaymentStrategy(paymentType string) (PaymentStrategy, error) {
	switch paymentType {
	case WeChatPayType:
		return &WeChatPay{}, nil
	case AliPayType:
		return &AliPay{}, nil
	case BankCardType:
		return &BankCard{}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}
