// 自动生成模板Setting
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Setting struct {
	gorm.Model
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
	Value   string `json:"value" form:"value" gorm:"column:value;comment:;type:varchar(255);size:255;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:;type:text;"`
}

func (Setting) TableName() string {
	return "setting"
}
