// 自动生成模板Link
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Link struct {
      gorm.Model
      ImgPath  string `json:"imgPath" form:"imgPath" gorm:"column:img_path;comment:;type:varchar(255);size:255;"`
      IsPub  int `json:"isPub" form:"isPub" gorm:"column:is_pub;comment:是否发布;type:int(10);size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;type:varchar(255);size:255;"`
      Sort  int `json:"sort" form:"sort" gorm:"column:sort;comment:;type:int(10);size:10;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:;type:int(10);size:10;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:;type:varchar(255);size:255;"` 
}


func (Link) TableName() string {
  return "link"
}
