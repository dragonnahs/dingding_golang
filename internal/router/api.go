package router

import (
	"github.com/gin-gonic/gin"
	"dingding_golang/internal/controller"
)

func RegisterAPIRoutes(r *gin.RouterGroup) {
	// 钉钉相关路由
	dingtalk := r.Group("/dingtalk")
	{
		c := controller.NewDingTalkController()
		dingtalk.POST("/webhook", c.Webhook)
		dingtalk.GET("/departments", c.GetDepartments)
		dingtalk.GET("/users", c.GetUsers)
	}
} 