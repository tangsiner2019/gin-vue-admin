// 自动生成模板Banner
package model

import (
      "gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Banner struct {
      gorm.Model
      ImgPath  string `json:"imgPath" form:"imgPath" gorm:"column:img_path;comment:;type:varchar(255);size:255;"`
      Link  string `json:"link" form:"link" gorm:"column:link;comment:;type:varchar(255);size:255;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
      IsPub  int `json:"isPub" form:"isPub" gorm:"column:is_pub;comment:;type:int(10);size:10;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(255);size:255;"`
}


func (Banner) TableName() string {
  return "banner"
}
