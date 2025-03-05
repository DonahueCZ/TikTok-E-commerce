package models

import "time"

type PaymentOrder struct {
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	OrderID       string    `gorm:"type:varchar(64);not null"`
	UserID        string    `gorm:"type:varchar(64);not null"` // ⚠️ **新增 UserID**
	Amount        float64   `gorm:"not null"`
	PaymentMethod string    `gorm:"type:varchar(20)"`
	Status        string    `gorm:"type:varchar(20);default:'pending'"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
