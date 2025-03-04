package models

import "time"

// Payment 代表支付记录
type Payment struct {
	PaymentID string  `gorm:"primaryKey"`
	OrderID   string  `gorm:"index"` // 关联订单
	Amount    float64 // 支付金额
	Status    string  // 支付状态: "Pending", "Paid", "Cancelled", "Timeout"
	CreatedAt time.Time
	UpdatedAt time.Time
}
