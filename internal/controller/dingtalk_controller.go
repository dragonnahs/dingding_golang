/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 13:53:58
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 15:30:21
 * @FilePath: \dingding_golang\internal\controller\dingtalk_controller.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"dingding_golang/internal/model"
	"dingding_golang/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DingTalkController struct {
	service *service.DingTalkService
}

func NewDingTalkController() *DingTalkController {
	return &DingTalkController{
		service: service.NewDingTalkService(),
	}
}

func (c *DingTalkController) Webhook(ctx *gin.Context) {
	// 处理钉钉回调
}

func (c *DingTalkController) GetDepartments(ctx *gin.Context) {
	// 获取部门列表
}

func (c *DingTalkController) GetUsers(ctx *gin.Context) {
	// 获取用户列表
}

// 发送工单卡片
func (c *DingTalkController) SendWorkOrderCard(ctx *gin.Context) {
	var order model.WorkOrder
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "无效的请求参数",
		})
		return
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "缺少用户ID",
		})
		return
	}

	userIdStr, ok := userId.(string)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "无效的用户ID类型",
		})
		return
	}

	if err := c.service.SendWorkOrderCard(userIdStr, &order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": "发送工单卡片失败",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "发送成功",
	})
}

// 创建开工单
func (c *DingTalkController) CreateWorkOrderApproval(ctx *gin.Context) {
	var order model.WorkOrder
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"message": "无效的请求参数",
		})
		return
	}

	if err := c.service.CreateWorkOrderApproval(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"message": "创建开工单失败",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "创建成功",
	})
} 