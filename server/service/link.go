package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateLink
// @description   create a Link
// @param     link               model.Link
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateLink(link model.Link) (err error) {
	err = global.GVA_DB.Create(&link).Error
	return err
}

// @title    DeleteLink
// @description   delete a Link
// @auth                     （2020/04/05  20:22）
// @param     link               model.Link
// @return                    error

func DeleteLink(link model.Link) (err error) {
	err = global.GVA_DB.Delete(link).Error
	return err
}

// @title    DeleteLinkByIds
// @description   delete Links
// @auth                     （2020/04/05  20:22）
// @param     link               model.Link
// @return                    error

func DeleteLinkByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Link{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateLink
// @description   update a Link
// @param     link          *model.Link
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateLink(link *model.Link) (err error) {
	err = global.GVA_DB.Save(link).Error
	return err
}

// @title    GetLink
// @description   get the info of a Link
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Link        Link

func GetLink(id uint) (err error, link model.Link) {
	err = global.GVA_DB.Where("id = ?", id).First(&link).Error
	return
}

// @title    GetLinkInfoList
// @description   get Link list by pagination, 分页获取Link
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetLinkInfoList(info request.LinkSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Link{})
	var links []model.Link
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&links).Error
	return err, links, total
}
