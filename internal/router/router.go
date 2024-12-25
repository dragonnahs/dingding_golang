package router

import (
	"dingding_golang/internal/middleware"
	"dingding_golang/pkg/config"

	"github.com/gin-gonic/gin"

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