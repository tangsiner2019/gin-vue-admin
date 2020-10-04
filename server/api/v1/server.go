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

// @Tags Server
// @Summary 创建Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Server true "创建Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/createServer [post]
func CreateServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindJSON(&server)
	err := service.CreateServer(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Server
// @Summary 删除Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Server true "删除Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /server/deleteServer [delete]
func DeleteServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindJSON(&server)
	err := service.DeleteServer(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Server
// @Summary 批量删除Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /server/deleteServerByIds [delete]
func DeleteServerByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteServerByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Server
// @Summary 更新Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Server true "更新Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /server/updateServer [put]
func UpdateServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindJSON(&server)
	err := service.UpdateServer(&server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Server
// @Summary 用id查询Server
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Server true "用id查询Server"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /server/findServer [get]
func FindServer(c *gin.Context) {
	var server model.Server
	_ = c.ShouldBindQuery(&server)
	err, reserver := service.GetServer(server.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reserver": reserver}, c)
	}
}

// @Tags Server
// @Summary 分页获取Server列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ServerSearch true "分页获取Server列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /server/getServerList [get]
func GetServerList(c *gin.Context) {
	var pageInfo request.ServerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetServerInfoList(pageInfo)
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
