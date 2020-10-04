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

// @Tags News
// @Summary 创建News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "创建News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/createNews [post]
func CreateNews(c *gin.Context) {
	var news model.News
	_ = c.ShouldBindJSON(&news)
	err := service.CreateNews(news)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags News
// @Summary 删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNews [delete]
func DeleteNews(c *gin.Context) {
	var news model.News
	_ = c.ShouldBindJSON(&news)
	err := service.DeleteNews(news)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags News
// @Summary 批量删除News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /news/deleteNewsByIds [delete]
func DeleteNewsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteNewsByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags News
// @Summary 更新News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "更新News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /news/updateNews [put]
func UpdateNews(c *gin.Context) {
	var news model.News
	_ = c.ShouldBindJSON(&news)
	err := service.UpdateNews(&news)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags News
// @Summary 用id查询News
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.News true "用id查询News"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /news/findNews [get]
func FindNews(c *gin.Context) {
	var news model.News
	_ = c.ShouldBindQuery(&news)
	err, renews := service.GetNews(news.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"renews": renews}, c)
	}
}

// @Tags News
// @Summary 分页获取News列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NewsSearch true "分页获取News列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /news/getNewsList [get]
func GetNewsList(c *gin.Context) {
	var pageInfo request.NewsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetNewsInfoList(pageInfo)
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
