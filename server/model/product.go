// 自动生成模板Product
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Product struct {
	gorm.Model
	CateId   int    `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:;type:int(10);size:10;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(255);size:255;"`
	SubTitle string `json:"subTitle" form:"subTitle" gorm:"column:sub_title;comment:;type:varchar(255);size:255;"`
	ImgPath  string `json:"imgPath" form:"imgPath" gorm:"column:img_path;comment:;type:varchar(255);size:255;"`
	Imgs     string `json:"imgs" form:"imgs" gorm:"column:imgs;comment:详情图json;type:text;"`
	Content  string `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`
	Tags     string `json:"tags" form:"tags" gorm:"column:tags;comment:标签json;type:varchar(255);size:255;"`
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
	Sort     int    `json:"sort" form:"sort" gorm:"column:sort;comment:;type:int(10);size:10;"`
	IsPub    int    `json:"isPub" form:"isPub" gorm:"column:is_pub;comment:是否发布;type:int(10);size:10;"`
}

func (Product) TableName() string {
	return "product"
}
