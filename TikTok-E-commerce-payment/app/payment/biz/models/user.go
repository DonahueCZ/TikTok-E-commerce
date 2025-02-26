package models

// User 表示用户模型
type User struct {
	ID      string  `gorm:"primaryKey" json:"id"` // 用户ID
	Balance float64 `json:"balance"`              // 用户余额
	Name    string  `json:"name"`                 // 用户姓名
	Email   string  `json:"email"`                // 用户邮箱
}
