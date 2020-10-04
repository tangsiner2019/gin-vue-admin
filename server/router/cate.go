package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCateRouter(Router *gin.RouterGroup) {
	CateRouter := Router.Group("cate").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		CateRouter.POST("createCate", v1.CreateCate)             // 新建Cate
		CateRouter.DELETE("deleteCate", v1.DeleteCate)           // 删除Cate
		CateRouter.DELETE("deleteCateByIds", v1.DeleteCateByIds) // 批量删除Cate
		CateRouter.PUT("updateCate", v1.UpdateCate)              // 更新Cate
		CateRouter.GET("findCate", v1.FindCate)                  // 根据ID获取Cate
		CateRouter.GET("getCateList", v1.GetCateList)            // 获取Cate列表
	}
}
