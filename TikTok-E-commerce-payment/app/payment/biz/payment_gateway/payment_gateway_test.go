package payment_gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// **测试 AliPay 支付**
func TestAliPay_Pay(t *testing.T) {
	alipay := &AliPay{}
	req := PaymentRequest{
		OrderID:     "order_1001",
		Amount:      100.00,
		PaymentType: AliPayType,
	}

	resp, err := alipay.Pay(req)

	assert.NoError(t, err)
	assert.Contains(t, resp, "AliPay Payment for Order order_1001 successful")
}

// **测试 WeChatPay 支付**
func TestWeChatPay_Pay(t *testing.T) {
	wechatPay := &WeChatPay{}
	req := PaymentRequest{
		OrderID:     "order_1002",
		Amount:      200.00,
		PaymentType: WeChatPayType,
	}

	resp, err := wechatPay.Pay(req)

	assert.NoError(t, err)
	assert.Contains(t, resp, "WeChat Payment for Order order_1002 successful")
}

// **测试 BankCard 支付**
func TestBankCard_Pay(t *testing.T) {
	bankCard := &BankCard{}
	req := PaymentRequest{
		OrderID:     "order_1003",
		Amount:      300.00,
		PaymentType: BankCardType,
	}

	resp, err := bankCard.Pay(req)

	assert.NoError(t, err)
	assert.Contains(t, resp, "BankCard Payment for Order order_1003 successful")
}

// **测试 PaymentFactory 获取支付方式**
func TestPaymentFactory_GetPaymentStrategy(t *testing.T) {
	factory := NewPaymentFactory()

	// 测试获取 AliPay
	strategy, err := factory.GetPaymentStrategy(AliPayType)
	assert.NoError(t, err)
	assert.IsType(t, &AliPay{}, strategy)

	// 测试获取 WeChatPay
	strategy, err = factory.GetPaymentStrategy(WeChatPayType)
	assert.NoError(t, err)
	assert.IsType(t, &WeChatPay{}, strategy)

	// 测试获取 BankCard
	strategy, err = factory.GetPaymentStrategy(BankCardType)
	assert.NoError(t, err)
	assert.IsType(t, &BankCard{}, strategy)

	// 测试获取错误的支付方式
	strategy, err = factory.GetPaymentStrategy("unknown")
	assert.Error(t, err)
	assert.Nil(t, strategy)
}
