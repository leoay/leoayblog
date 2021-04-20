package models

import (
	"time"
)

//User 用户模型
type User struct {
	ID        int64     `gorm:"primaryKey"`
	Name      string    `gorm:"unique"`
	Password  string    `gorm:"size(100)"`
	Email     string    `gorm:"size(100)"`
	Sex       string    `gorm:"size(2)"`
	Phone     string    `orm:"size(16)"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
	UpdatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
}

//TableName 指定数据库名称
func (User) TableName() string {
	return "users"
}
