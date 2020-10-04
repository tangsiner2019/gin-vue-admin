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

// @Tags Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/createProduct [post]
func CreateProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	err := service.CreateProduct(product)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProduct [delete]
func DeleteProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	err := service.DeleteProduct(product)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Product
// @Summary 批量删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /product/deleteProductByIds [delete]
func DeleteProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	err := service.DeleteProductByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /product/updateProduct [put]
func UpdateProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindJSON(&product)
	err := service.UpdateProduct(&product)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /product/findProduct [get]
func FindProduct(c *gin.Context) {
	var product model.Product
	_ = c.ShouldBindQuery(&product)
	err, reproduct := service.GetProduct(product.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reproduct": reproduct}, c)
	}
}

// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProductSearch true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /product/getProductList [get]
func GetProductList(c *gin.Context) {
	var pageInfo request.ProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetProductInfoList(pageInfo)
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
