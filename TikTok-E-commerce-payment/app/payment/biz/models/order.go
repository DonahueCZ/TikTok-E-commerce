package models

import (
	"time"
)

// PaymentOrder 定义数据库中的订单结构
type PaymentOrder struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`           // 主键，自动递增
	OrderID       string    `gorm:"type:varchar(64);not null"`          // 订单ID，最大64字符
	Amount        float64   `gorm:"not null"`                           // 金额，浮动
	Status        string    `gorm:"type:varchar(20);default:'pending'"` // 状态，默认为 'pending'
	PaymentMethod string    `gorm:"type:varchar(20)"`                   // 支付方式
	CreatedAt     time.Time `gorm:"autoCreateTime"`                     // 创建时间
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`                     // 更新时间
}
