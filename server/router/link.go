package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitLinkRouter(Router *gin.RouterGroup) {
	LinkRouter := Router.Group("link").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		LinkRouter.POST("createLink", v1.CreateLink)   // 新建Link
		LinkRouter.DELETE("deleteLink", v1.DeleteLink) // 删除Link
		LinkRouter.DELETE("deleteLinkByIds", v1.DeleteLinkByIds) // 批量删除Link
		LinkRouter.PUT("updateLink", v1.UpdateLink)    // 更新Link
		LinkRouter.GET("findLink", v1.FindLink)        // 根据ID获取Link
		LinkRouter.GET("getLinkList", v1.GetLinkList)  // 获取Link列表
	}
}
