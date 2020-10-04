package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTeamRouter(Router *gin.RouterGroup) {
	TeamRouter := Router.Group("team").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		TeamRouter.POST("createTeam", v1.CreateTeam)             // 新建Team
		TeamRouter.DELETE("deleteTeam", v1.DeleteTeam)           // 删除Team
		TeamRouter.DELETE("deleteTeamByIds", v1.DeleteTeamByIds) // 批量删除Team
		TeamRouter.PUT("updateTeam", v1.UpdateTeam)              // 更新Team
		TeamRouter.GET("findTeam", v1.FindTeam)                  // 根据ID获取Team
		TeamRouter.GET("getTeamList", v1.GetTeamList)            // 获取Team列表
	}
}
