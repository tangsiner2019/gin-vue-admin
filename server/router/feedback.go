package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitFeedbackRouter(Router *gin.RouterGroup) {
	FeedbackRouter := Router.Group("feedback").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		FeedbackRouter.POST("createFeedback", v1.CreateFeedback)             // 新建Feedback
		FeedbackRouter.DELETE("deleteFeedback", v1.DeleteFeedback)           // 删除Feedback
		FeedbackRouter.DELETE("deleteFeedbackByIds", v1.DeleteFeedbackByIds) // 批量删除Feedback
		FeedbackRouter.PUT("updateFeedback", v1.UpdateFeedback)              // 更新Feedback
		FeedbackRouter.GET("findFeedback", v1.FindFeedback)                  // 根据ID获取Feedback
		FeedbackRouter.GET("getFeedbackList", v1.GetFeedbackList)            // 获取Feedback列表
	}
}
