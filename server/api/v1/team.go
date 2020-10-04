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

// @Tags Team
// @Summary 创建Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Team true "创建Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/createTeam [post]
func CreateTeam(c *gin.Context) {
	var team model.Team
	_ = c.ShouldBindJSON(&team)
	err := service.CreateTeam(team)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Team
// @Summary 删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Team true "删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeam [delete]
func DeleteTeam(c *gin.Context) {
	var team model.Team
	_ = c.ShouldBindJSON(&team)
	err := service.DeleteTeam(team)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Team
// @Summary 批量删除Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /team/deleteTeamByIds [delete]
func DeleteTeamByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteTeamByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Team
// @Summary 更新Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Team true "更新Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /team/updateTeam [put]
func UpdateTeam(c *gin.Context) {
	var team model.Team
	_ = c.ShouldBindJSON(&team)
	err := service.UpdateTeam(&team)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Team
// @Summary 用id查询Team
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Team true "用id查询Team"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /team/findTeam [get]
func FindTeam(c *gin.Context) {
	var team model.Team
	_ = c.ShouldBindQuery(&team)
	err, reteam := service.GetTeam(team.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reteam": reteam}, c)
	}
}

// @Tags Team
// @Summary 分页获取Team列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TeamSearch true "分页获取Team列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /team/getTeamList [get]
func GetTeamList(c *gin.Context) {
	var pageInfo request.TeamSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetTeamInfoList(pageInfo)
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
