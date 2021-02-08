package models

import (
	_ "gorm.io/gorm"
)

type Create_daliy_thinks struct {
	Id          int
	Label       string
	Content     string
	ApprovalNum int
	StarNum     int
	Content_md  string
	Comments    string
	DateTime    string
}

func (Create_daliy_thinks) TableName() string {
	return "create_daliy_thinks"
}
