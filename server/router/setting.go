package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSettingRouter(Router *gin.RouterGroup) {
	SettingRouter := Router.Group("setting").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		SettingRouter.POST("createSetting", v1.CreateSetting)             // 新建Setting
		SettingRouter.DELETE("deleteSetting", v1.DeleteSetting)           // 删除Setting
		SettingRouter.DELETE("deleteSettingByIds", v1.DeleteSettingByIds) // 批量删除Setting
		SettingRouter.PUT("updateSetting", v1.UpdateSetting)              // 更新Setting
		SettingRouter.GET("findSetting", v1.FindSetting)                  // 根据ID获取Setting
		SettingRouter.GET("getSettingList", v1.GetSettingList)            // 获取Setting列表
	}
}
