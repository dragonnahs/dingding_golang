package router

import (
	"github.com/gin-gonic/gin"
	"your-project/internal/controller"
)

func RegisterAPIRoutes(r *gin.RouterGroup) {
	// 钉钉相关路由
	dingtalk := r.Group("/dingtalk")
	{
		c := controller.NewDingtalkController()
		dingtalk.POST("/webhook", c.Webhook)
		dingtalk.GET("/departments", c.GetDepartments)
		dingtalk.GET("/users", c.GetUsers)
	}
} 