/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 13:53:58
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 14:16:46
 * @FilePath: \dingding_golang\internal\controller\dingtalk_controller.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controller

import (
	"dingding_golang/internal/service"

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