package models

import (
	_ "gorm.io/gorm"
)

//文章类型模型
type Article_classes struct {
	Id               uint `gorm:"primary_key"`
	FirstClass       string
	SecondClass      string
	FirstClassRoute  string
	SecondClassRoute string
}

func (Article_classes) TableName() string {
	return "article_classes"
}
