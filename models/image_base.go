package models

import (
	_ "gorm.io/gorm"
)

//文章类型模型
type Image_Base struct {
	Id        uint `gorm:"primary_key"`
	Name      string
	Path      string
	UserId    int64
	CreatedAt uint64
	UpdatedAt uint64
}

func (Image_Base) TableName() string {
	return "image_bases"
}
