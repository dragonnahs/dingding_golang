package router

import (
	"github.com/gin-gonic/gin"
	"your-project/internal/middleware"
	"your-project/pkg/config"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Get().Server.Mode)
	
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	
	// API 路由组
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.Auth())
	
	// 注册 API 路由
	RegisterAPIRoutes(apiGroup)
	
	return r
} 