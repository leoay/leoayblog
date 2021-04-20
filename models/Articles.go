package models

import "time"

//Articles 文章基础模型
type Articles struct {
	ID          int64 `gorm:"primary_key"`
	UserID      int64
	Author      string
	Title       string `gorm:"type:longText"`
	MainPicPATH string
	Summary     string
	ClassID     int64
	Content     string `gorm:"type:longText"`
	ReadNum     int
	ApprovalNum int
	StarNum     int
	IsTemp      bool
	ContentMd   string    `gorm:"type:longText"`
	Comments    string    `gorm:"type:longText"`
	CreatedAt   time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
	UpdatedAt   time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
}

//TableName 数据表名
func (Articles) TableName() string {
	return "articles"
}
