package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateBanner
// @description   create a Banner
// @param     banner               model.Banner
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateBanner(banner model.Banner) (err error) {
	err = global.GVA_DB.Create(&banner).Error
	return err
}

// @title    DeleteBanner
// @description   delete a Banner
// @auth                     （2020/04/05  20:22）
// @param     banner               model.Banner
// @return                    error

func DeleteBanner(banner model.Banner) (err error) {
	err = global.GVA_DB.Delete(banner).Error
	return err
}

// @title    DeleteBannerByIds
// @description   delete Banners
// @auth                     （2020/04/05  20:22）
// @param     banner               model.Banner
// @return                    error

func DeleteBannerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Banner{},"id in ?",ids.Ids).Error
	return err
}

// @title    UpdateBanner
// @description   update a Banner
// @param     banner          *model.Banner
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateBanner(banner *model.Banner) (err error) {
	err = global.GVA_DB.Save(banner).Error
	return err
}

// @title    GetBanner
// @description   get the info of a Banner
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Banner        Banner

func GetBanner(id uint) (err error, banner model.Banner) {
	err = global.GVA_DB.Where("id = ?", id).First(&banner).Error
	return
}

// @title    GetBannerInfoList
// @description   get Banner list by pagination, 分页获取Banner
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetBannerInfoList(info request.BannerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Banner{})
    var banners []model.Banner
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&banners).Error
	return err, banners, total
}