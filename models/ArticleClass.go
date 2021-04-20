package models

import (
	"time"
)

//ArticleClass 文章类型模型
type ArticleClass struct {
	ID               int64 `gorm:"primary_key"`
	FirstClass       string
	SecondClass      string
	FirstClassRoute  string
	SecondClassRoute string
	CreatedAt        time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
	UpdatedAt        time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
}

//TableName 数据表
func (ArticleClass) TableName() string {
	return "article_class"
}
