package router

import (
	"dingding_golang/internal/controller"
	"dingding_golang/internal/middleware"

	"github.com/gin-gonic/gin"

)

func RegisterAPIRoutes(r *gin.RouterGroup) {
	// 认证路由 - 不需要验证token
	auth := r.Group("/auth")
	{
		c := controller.NewAuthController()
		auth.POST("/token", c.GenerateToken)
	}

	// 钉钉相关路由 - 需要验证token
	dingtalk := r.Group("/dingtalk")
	dingtalk.Use(middleware.Auth())
	{
		c := controller.NewDingTalkController()
		dingtalk.POST("/webhook", c.Webhook)
		dingtalk.GET("/departments", c.GetDepartments)
		dingtalk.GET("/users", c.GetUsers)
		dingtalk.POST("/workorder/card", c.SendWorkOrderCard)
		dingtalk.POST("/workorder/approval", c.CreateWorkOrderApproval)
	}
} 