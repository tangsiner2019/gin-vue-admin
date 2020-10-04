// 自动生成模板Cate
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Cate struct {
	gorm.Model
	Name   string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
	Sort   int    `json:"sort" form:"sort" gorm:"column:sort;comment:;type:int(10);size:10;"`
	Status int    `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
	Type   string `json:"type" form:"type" gorm:"column:type;comment:;type:varchar(255);size:255;"`
}

func (Cate) TableName() string {
	return "cate"
}
