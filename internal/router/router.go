/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 13:53:52
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 15:11:57
 * @FilePath: \dingding_golang\internal\router\router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
	
	// 注册 API 路由
	RegisterAPIRoutes(apiGroup)
	
	return r
} 