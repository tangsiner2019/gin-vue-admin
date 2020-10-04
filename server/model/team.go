// 自动生成模板Team
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Team struct {
	gorm.Model
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
	ImgPath string `json:"imgPath" form:"imgPath" gorm:"column:img_path;comment:;type:varchar(255);size:255;"`
	EnName  string `json:"enName" form:"enName" gorm:"column:en_name;comment:;type:varchar(255);size:255;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
	Sort    int    `json:"sort" form:"sort" gorm:"column:sort;comment:;type:int(10);size:10;"`
	IsPub   int    `json:"isPub" form:"isPub" gorm:"column:is_pub;comment:是否发布;type:int(10);size:10;"`
}

func (Team) TableName() string {
	return "team"
}
