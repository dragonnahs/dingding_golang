/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 14:30:00
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 18:17:39
 * @FilePath: \dingding_golang\internal\service\dingtalk_service.go
 */
package service

import (
	"dingding_golang/internal/model"
	"fmt"

	"dingding_golang/pkg/dingtalk"
)

type DingTalkService struct {
	client *dingtalk.DingTalkClient
}

func NewDingTalkService() *DingTalkService {
	return &DingTalkService{
		client: dingtalk.GetClient(),
	}
}

// 发送工单卡片
func (s *DingTalkService) SendWorkOrderCard(userId string, order *model.WorkOrder) error {
	cardData := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title": order.Title,
			"text": fmt.Sprintf("工单号: %s\n描述: %s\n地址: %s", 
				order.ID, order.Description, order.Location.Address),
			"btnOrientation": "0",
			"btns": []map[string]string{
				{
					"title": "创建开工单",
					"actionURL": fmt.Sprintf("dingtalk://dingtalkclient/page/link?url=%s&pc_slide=false",
						s.generateApprovalUrl(order)),
				},
			},
		},
	}

	// 发送消息
	return s.sendMessage(userId, cardData)
}

// 创建审批实例
func (s *DingTalkService) CreateWorkOrderApproval(order *model.WorkOrder) error {
	// 构建审批表单数据
	formData := []map[string]interface{}{
		{
			"name": "工单号",
			"value": order.ID,
		},
		{
			"name": "工单标题",
			"value": order.Title,
		},
		{
			"name": "工单描述",
			"value": order.Description,
		},
		{
			"name": "定位信息",
			"value": map[string]interface{}{
				"latitude": order.Location.Latitude,
				"longitude": order.Location.Longitude,
				"address": order.Location.Address,
			},
		},
	}

	// 如果有图片，添加图片字段
	if len(order.Images) > 0 {
		formData = append(formData, map[string]interface{}{
			"name": "现场图片",
			"value": order.Images,
		})
	}

	approvalData := map[string]interface{}{
		"process_code": "PROC-XXX", // 审批流程码
		"originator_user_id": order.CreatorId,
		"form_component_values": formData,
	}

	return s.createApproval(approvalData)
}

// 生成审批链接
func (s *DingTalkService) generateApprovalUrl(order *model.WorkOrder) string {
	// 生成审批页面URL，包含工单信息
	return fmt.Sprintf("https://your-domain.com/approval?orderId=%s", order.ID)
}

// 发送消息到钉钉
func (s *DingTalkService) sendMessage(userId string, message interface{}) error {
	return s.client.SendMessage(userId, message)
}

// 创建审批实例
func (s *DingTalkService) createApproval(approvalData interface{}) error {
	// 调用钉钉API创建审批实例
	// TODO: 实现具体的审批创建逻辑

	return nil
} 