package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateSetting
// @description   create a Setting
// @param     setting               model.Setting
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateSetting(setting model.Setting) (err error) {
	err = global.GVA_DB.Create(&setting).Error
	return err
}

// @title    DeleteSetting
// @description   delete a Setting
// @auth                     （2020/04/05  20:22）
// @param     setting               model.Setting
// @return                    error

func DeleteSetting(setting model.Setting) (err error) {
	err = global.GVA_DB.Delete(setting).Error
	return err
}

// @title    DeleteSettingByIds
// @description   delete Settings
// @auth                     （2020/04/05  20:22）
// @param     setting               model.Setting
// @return                    error

func DeleteSettingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Setting{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateSetting
// @description   update a Setting
// @param     setting          *model.Setting
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateSetting(setting *model.Setting) (err error) {
	err = global.GVA_DB.Save(setting).Error
	return err
}

// @title    GetSetting
// @description   get the info of a Setting
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Setting        Setting

func GetSetting(id uint) (err error, setting model.Setting) {
	err = global.GVA_DB.Where("id = ?", id).First(&setting).Error
	return
}

// @title    GetSettingInfoList
// @description   get Setting list by pagination, 分页获取Setting
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetSettingInfoList(info request.SettingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Setting{})
	var settings []model.Setting
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&settings).Error
	return err, settings, total
}
