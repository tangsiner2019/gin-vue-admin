package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitNewsRouter(Router *gin.RouterGroup) {
	NewsRouter := Router.Group("news").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		NewsRouter.POST("createNews", v1.CreateNews)             // 新建News
		NewsRouter.DELETE("deleteNews", v1.DeleteNews)           // 删除News
		NewsRouter.DELETE("deleteNewsByIds", v1.DeleteNewsByIds) // 批量删除News
		NewsRouter.PUT("updateNews", v1.UpdateNews)              // 更新News
		NewsRouter.GET("findNews", v1.FindNews)                  // 根据ID获取News
		NewsRouter.GET("getNewsList", v1.GetNewsList)            // 获取News列表
	}
}
