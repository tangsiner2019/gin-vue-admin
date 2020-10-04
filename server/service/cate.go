package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateCate
// @description   create a Cate
// @param     cate               model.Cate
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateCate(cate model.Cate) (err error) {
	err = global.GVA_DB.Create(&cate).Error
	return err
}

// @title    DeleteCate
// @description   delete a Cate
// @auth                     （2020/04/05  20:22）
// @param     cate               model.Cate
// @return                    error

func DeleteCate(cate model.Cate) (err error) {
	err = global.GVA_DB.Delete(cate).Error
	return err
}

// @title    DeleteCateByIds
// @description   delete Cates
// @auth                     （2020/04/05  20:22）
// @param     cate               model.Cate
// @return                    error

func DeleteCateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Cate{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateCate
// @description   update a Cate
// @param     cate          *model.Cate
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateCate(cate *model.Cate) (err error) {
	err = global.GVA_DB.Save(cate).Error
	return err
}

// @title    GetCate
// @description   get the info of a Cate
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Cate        Cate

func GetCate(id uint) (err error, cate model.Cate) {
	err = global.GVA_DB.Where("id = ?", id).First(&cate).Error
	return
}

// @title    GetCateInfoList
// @description   get Cate list by pagination, 分页获取Cate
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetCateInfoList(info request.CateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Cate{})
	var cates []model.Cate
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&cates).Error
	return err, cates, total
}
