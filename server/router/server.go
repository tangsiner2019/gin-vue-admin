package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(Router *gin.RouterGroup) {
	ServerRouter := Router.Group("server").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ServerRouter.POST("createServer", v1.CreateServer)             // 新建Server
		ServerRouter.DELETE("deleteServer", v1.DeleteServer)           // 删除Server
		ServerRouter.DELETE("deleteServerByIds", v1.DeleteServerByIds) // 批量删除Server
		ServerRouter.PUT("updateServer", v1.UpdateServer)              // 更新Server
		ServerRouter.GET("findServer", v1.FindServer)                  // 根据ID获取Server
		ServerRouter.GET("getServerList", v1.GetServerList)            // 获取Server列表
	}
}
