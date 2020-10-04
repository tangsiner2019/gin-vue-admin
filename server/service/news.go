package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateNews
// @description   create a News
// @param     news               model.News
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateNews(news model.News) (err error) {
	err = global.GVA_DB.Create(&news).Error
	return err
}

// @title    DeleteNews
// @description   delete a News
// @auth                     （2020/04/05  20:22）
// @param     news               model.News
// @return                    error

func DeleteNews(news model.News) (err error) {
	err = global.GVA_DB.Delete(news).Error
	return err
}

// @title    DeleteNewsByIds
// @description   delete Newss
// @auth                     （2020/04/05  20:22）
// @param     news               model.News
// @return                    error

func DeleteNewsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.News{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateNews
// @description   update a News
// @param     news          *model.News
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateNews(news *model.News) (err error) {
	err = global.GVA_DB.Save(news).Error
	return err
}

// @title    GetNews
// @description   get the info of a News
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    News        News

func GetNews(id uint) (err error, news model.News) {
	err = global.GVA_DB.Where("id = ?", id).First(&news).Error
	return
}

// @title    GetNewsInfoList
// @description   get News list by pagination, 分页获取News
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetNewsInfoList(info request.NewsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.News{})
	var newss []model.News
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&newss).Error
	return err, newss, total
}
