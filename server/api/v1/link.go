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

// @Tags Link
// @Summary 创建Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "创建Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /link/createLink [post]
func CreateLink(c *gin.Context) {
	var link model.Link
	_ = c.ShouldBindJSON(&link)
	err := service.CreateLink(link)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Link
// @Summary 删除Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "删除Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /link/deleteLink [delete]
func DeleteLink(c *gin.Context) {
	var link model.Link
	_ = c.ShouldBindJSON(&link)
	err := service.DeleteLink(link)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Link
// @Summary 批量删除Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /link/deleteLinkByIds [delete]
func DeleteLinkByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteLinkByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Link
// @Summary 更新Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "更新Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /link/updateLink [put]
func UpdateLink(c *gin.Context) {
	var link model.Link
	_ = c.ShouldBindJSON(&link)
	err := service.UpdateLink(&link)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Link
// @Summary 用id查询Link
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Link true "用id查询Link"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /link/findLink [get]
func FindLink(c *gin.Context) {
	var link model.Link
	_ = c.ShouldBindQuery(&link)
	err, relink := service.GetLink(link.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"relink": relink}, c)
	}
}

// @Tags Link
// @Summary 分页获取Link列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LinkSearch true "分页获取Link列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /link/getLinkList [get]
func GetLinkList(c *gin.Context) {
	var pageInfo request.LinkSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetLinkInfoList(pageInfo)
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
