// 自动生成模板Feedback
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Feedback struct {
	gorm.Model
	Name    string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
	Phone   string `json:"phone" form:"phone" gorm:"column:phone;comment:;type:varchar(255);size:255;"`
	Email   string `json:"email" form:"email" gorm:"column:email;comment:;type:varchar(255);size:255;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:;type:varchar(255);size:255;"`
	Status  int    `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
}

func (Feedback) TableName() string {
	return "feedback"
}
