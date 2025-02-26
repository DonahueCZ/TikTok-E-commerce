package models

// PaymentDetails 表示支付相关的细节
type PaymentDetails struct {
	ProductID string  `json:"product_id"` // 产品ID
	UserID    string  `json:"user_id"`    // 用户ID
	Quantity  int     `json:"quantity"`   // 产品数量
	Amount    float64 `json:"amount"`     // 支付金额
}

// Product 表示一个商品
type Product struct {
	ID    string `json:"id" gorm:"primaryKey"` // 产品ID
	Stock int    `json:"stock"`                // 库存数量
}
