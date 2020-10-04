package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("product").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ProductRouter.POST("createProduct", v1.CreateProduct)             // 新建Product
		ProductRouter.DELETE("deleteProduct", v1.DeleteProduct)           // 删除Product
		ProductRouter.DELETE("deleteProductByIds", v1.DeleteProductByIds) // 批量删除Product
		ProductRouter.PUT("updateProduct", v1.UpdateProduct)              // 更新Product
		ProductRouter.GET("findProduct", v1.FindProduct)                  // 根据ID获取Product
		ProductRouter.GET("getProductList", v1.GetProductList)            // 获取Product列表
	}
}
