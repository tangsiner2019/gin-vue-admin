package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Setting
// @Summary 创建Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "创建Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/createSetting [post]
func CreateSetting(c *gin.Context) {
	var setting model.Setting
	_ = c.ShouldBindJSON(&setting)
	err := service.CreateSetting(setting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Setting
// @Summary 删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /setting/deleteSetting [delete]
func DeleteSetting(c *gin.Context) {
	var setting model.Setting
	_ = c.ShouldBindJSON(&setting)
	err := service.DeleteSetting(setting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Setting
// @Summary 批量删除Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /setting/deleteSettingByIds [delete]
func DeleteSettingByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteSettingByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Setting
// @Summary 更新Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "更新Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /setting/updateSetting [put]
func UpdateSetting(c *gin.Context) {
	var setting model.Setting
	_ = c.ShouldBindJSON(&setting)
	err := service.UpdateSetting(&setting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Setting
// @Summary 用id查询Setting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Setting true "用id查询Setting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /setting/findSetting [get]
func FindSetting(c *gin.Context) {
	var setting model.Setting
	_ = c.ShouldBindQuery(&setting)
	err, resetting := service.GetSetting(setting.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"resetting": resetting}, c)
	}
}

// @Tags Setting
// @Summary 分页获取Setting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SettingSearch true "分页获取Setting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /setting/getSettingList [get]
func GetSettingList(c *gin.Context) {
	var pageInfo request.SettingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetSettingInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
