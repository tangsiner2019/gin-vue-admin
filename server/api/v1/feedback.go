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

// @Tags Feedback
// @Summary 创建Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Feedback true "创建Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /feedback/createFeedback [post]
func CreateFeedback(c *gin.Context) {
	var feedback model.Feedback
	_ = c.ShouldBindJSON(&feedback)
	err := service.CreateFeedback(feedback)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Feedback
// @Summary 删除Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Feedback true "删除Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /feedback/deleteFeedback [delete]
func DeleteFeedback(c *gin.Context) {
	var feedback model.Feedback
	_ = c.ShouldBindJSON(&feedback)
	err := service.DeleteFeedback(feedback)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Feedback
// @Summary 批量删除Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /feedback/deleteFeedbackByIds [delete]
func DeleteFeedbackByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteFeedbackByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Feedback
// @Summary 更新Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Feedback true "更新Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /feedback/updateFeedback [put]
func UpdateFeedback(c *gin.Context) {
	var feedback model.Feedback
	_ = c.ShouldBindJSON(&feedback)
	err := service.UpdateFeedback(&feedback)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Feedback
// @Summary 用id查询Feedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Feedback true "用id查询Feedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /feedback/findFeedback [get]
func FindFeedback(c *gin.Context) {
	var feedback model.Feedback
	_ = c.ShouldBindQuery(&feedback)
	err, refeedback := service.GetFeedback(feedback.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"refeedback": refeedback}, c)
	}
}

// @Tags Feedback
// @Summary 分页获取Feedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FeedbackSearch true "分页获取Feedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /feedback/getFeedbackList [get]
func GetFeedbackList(c *gin.Context) {
	var pageInfo request.FeedbackSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetFeedbackInfoList(pageInfo)
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
