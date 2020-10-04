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

// @Tags Cate
// @Summary 创建Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "创建Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cate/createCate [post]
func CreateCate(c *gin.Context) {
	var cate model.Cate
	_ = c.ShouldBindJSON(&cate)
	err := service.CreateCate(cate)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Cate
// @Summary 删除Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "删除Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cate/deleteCate [delete]
func DeleteCate(c *gin.Context) {
	var cate model.Cate
	_ = c.ShouldBindJSON(&cate)
	err := service.DeleteCate(cate)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Cate
// @Summary 批量删除Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cate/deleteCateByIds [delete]
func DeleteCateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteCateByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Cate
// @Summary 更新Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "更新Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cate/updateCate [put]
func UpdateCate(c *gin.Context) {
	var cate model.Cate
	_ = c.ShouldBindJSON(&cate)
	err := service.UpdateCate(&cate)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Cate
// @Summary 用id查询Cate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Cate true "用id查询Cate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cate/findCate [get]
func FindCate(c *gin.Context) {
	var cate model.Cate
	_ = c.ShouldBindQuery(&cate)
	err, recate := service.GetCate(cate.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"recate": recate}, c)
	}
}

// @Tags Cate
// @Summary 分页获取Cate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CateSearch true "分页获取Cate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cate/getCateList [get]
func GetCateList(c *gin.Context) {
	var pageInfo request.CateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetCateInfoList(pageInfo)
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
