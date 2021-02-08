package models

import (
	_ "gorm.io/gorm"
)

//文章基础模型
type Article_bases struct {
	Id          uint `gorm:"primary_key"`
	UserId      int64
	Author      string
	Title       string `gorm:"type:longText"`
	EditorId    uint
	MainPicPath string
	Summary     string
	ClassId     uint
	Content     string `gorm:"type:longText"`
	ApprovalNum int
	StarNum     int
	IsTemp      bool
	Content_md  string `gorm:"type:longText"`
	Comments    string `gorm:"type:longText"`
	CreatedAt   uint64
	UpdatedAt   uint64
}

func (Article_bases) TableName() string {
	return "article_bases"
}
