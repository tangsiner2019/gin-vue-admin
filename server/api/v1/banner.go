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

// @Tags Banner
// @Summary 创建Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "创建Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/createBanner [post]
func CreateBanner(c *gin.Context) {
	var banner model.Banner
	_ = c.ShouldBindJSON(&banner)
	err := service.CreateBanner(banner)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Banner
// @Summary 删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBanner [delete]
func DeleteBanner(c *gin.Context) {
	var banner model.Banner
	_ = c.ShouldBindJSON(&banner)
	err := service.DeleteBanner(banner)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Banner
// @Summary 批量删除Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banner/deleteBannerByIds [delete]
func DeleteBannerByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteBannerByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Banner
// @Summary 更新Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "更新Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /banner/updateBanner [put]
func UpdateBanner(c *gin.Context) {
	var banner model.Banner
	_ = c.ShouldBindJSON(&banner)
	err := service.UpdateBanner(&banner)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Banner
// @Summary 用id查询Banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Banner true "用id查询Banner"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banner/findBanner [get]
func FindBanner(c *gin.Context) {
	var banner model.Banner
	_ = c.ShouldBindQuery(&banner)
	err, rebanner := service.GetBanner(banner.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"rebanner": rebanner}, c)
	}
}

// @Tags Banner
// @Summary 分页获取Banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BannerSearch true "分页获取Banner列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banner/getBannerList [get]
func GetBannerList(c *gin.Context) {
	var pageInfo request.BannerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetBannerInfoList(pageInfo)
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
