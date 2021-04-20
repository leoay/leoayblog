package models

import "time"

//Images 文章类型模型
type Images struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size(100)"`
	Path      string    `gorm:"size(100)"`
	UserID    int64     `gorm:"size(10)"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
	UpdatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
}

//TableName 数据表名
func (Images) TableName() string {
	return "images"
}
