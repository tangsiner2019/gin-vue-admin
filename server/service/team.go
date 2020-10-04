package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    CreateTeam
// @description   create a Team
// @param     team               model.Team
// @auth                     （2020/04/05  20:22）
// @return    err             error

func CreateTeam(team model.Team) (err error) {
	err = global.GVA_DB.Create(&team).Error
	return err
}

// @title    DeleteTeam
// @description   delete a Team
// @auth                     （2020/04/05  20:22）
// @param     team               model.Team
// @return                    error

func DeleteTeam(team model.Team) (err error) {
	err = global.GVA_DB.Delete(team).Error
	return err
}

// @title    DeleteTeamByIds
// @description   delete Teams
// @auth                     （2020/04/05  20:22）
// @param     team               model.Team
// @return                    error

func DeleteTeamByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Team{}, "id in ?", ids.Ids).Error
	return err
}

// @title    UpdateTeam
// @description   update a Team
// @param     team          *model.Team
// @auth                     （2020/04/05  20:22）
// @return                    error

func UpdateTeam(team *model.Team) (err error) {
	err = global.GVA_DB.Save(team).Error
	return err
}

// @title    GetTeam
// @description   get the info of a Team
// @auth                     （2020/04/05  20:22）
// @param     id              uint
// @return                    error
// @return    Team        Team

func GetTeam(id uint) (err error, team model.Team) {
	err = global.GVA_DB.Where("id = ?", id).First(&team).Error
	return
}

// @title    GetTeamInfoList
// @description   get Team list by pagination, 分页获取Team
// @auth                     （2020/04/05  20:22）
// @param     info            PageInfo
// @return                    error

func GetTeamInfoList(info request.TeamSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Team{})
	var teams []model.Team
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&teams).Error
	return err, teams, total
}
