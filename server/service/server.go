package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateServer
// @description   create a Server
// @param     server               model.Server
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateServer(server model.Server) (err error) {
	err = global.GVA_DB.Create(&server).Error
	return err
}

// @title    DeleteServer
// @description   delete a Server
// @auth                     （2020/04/05  20:22）
// @param     server               model.Server
// @return                    error

func DeleteServer(server model.Server) (err error) {
	err = global.GVA_DB.Delete(server).Error
	return err
}

// @title    DeleteServerByIds
// @description   delete Servers
// @auth                     （2020/04/05  20:22）
// @param     server               model.Server
// @return                    error

func DeleteServerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Server{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateServer
// @description   update a Server
// @param     server          *model.Server
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateServer(server *model.Server) (err error) {
	err = global.GVA_DB.Save(server).Error
	return err
}

// @title    GetServer
// @description   get the info of a Server
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Server        Server

func GetServer(id uint) (err error, server model.Server) {
	err = global.GVA_DB.Where("id = ?", id).First(&server).Error
	return
}

// @title    GetServerInfoList
// @description   get Server list by pagination, 分页获取Server
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetServerInfoList(info request.ServerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Server{})
	var servers []model.Server
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&servers).Error
	return err, servers, total
}
