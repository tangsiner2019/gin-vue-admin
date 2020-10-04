// 自动生成模板News
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type News struct {
	gorm.Model
	CateId   int    `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:;type:int(10);size:10;"`
	Content  string `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`
	ImgPath  string `json:"imgPath" form:"imgPath" gorm:"column:img_path;comment:;type:varchar(255);size:255;"`
	IsPub    int    `json:"isPub" form:"isPub" gorm:"column:is_pub;comment:是否发布;type:int(10);size:10;"`
	Link     string `json:"link" form:"link" gorm:"column:link;comment:;type:varchar(255);size:255;"`
	Sort     int    `json:"sort" form:"sort" gorm:"column:sort;comment:;type:int(10);size:10;"`
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
	SubTitle string `json:"subTitle" form:"subTitle" gorm:"column:sub_title;comment:;type:varchar(255);size:255;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(255);size:255;"`
}

func (News) TableName() string {
	return "news"
}
