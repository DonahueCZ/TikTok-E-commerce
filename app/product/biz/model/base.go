package model

import "time"

type Base struct {
	ID        uint32 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
