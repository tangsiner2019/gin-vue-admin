package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitBannerRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("banner").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		BannerRouter.POST("createBanner", v1.CreateBanner)   // 新建Banner
		BannerRouter.DELETE("deleteBanner", v1.DeleteBanner) // 删除Banner
		BannerRouter.DELETE("deleteBannerByIds", v1.DeleteBannerByIds) // 批量删除Banner
		BannerRouter.PUT("updateBanner", v1.UpdateBanner)    // 更新Banner
		BannerRouter.GET("findBanner", v1.FindBanner)        // 根据ID获取Banner
		BannerRouter.GET("getBannerList", v1.GetBannerList)  // 获取Banner列表
	}
}
