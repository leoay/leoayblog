package models

//图片轮播模型
import (
	_ "gorm.io/gorm"
)

type Image_flows struct {
	Id       int
	ImgUrl   string
	DateTime string
}

func (Image_flows) TableName() string {
	return "image_flows"
}