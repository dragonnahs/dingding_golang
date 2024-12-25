/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 14:30:00
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 14:30:00
 * @FilePath: \dingding_golang\internal\model\work_order.go
 */
package model

type WorkOrder struct {
    ID          string   `json:"id"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    Location    Location `json:"location"`
    Images      []string `json:"images"`
    CreatorId   string   `json:"creator_id"`
    Status      string   `json:"status"`
    CreateTime  int64    `json:"create_time"`
}

type Location struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Address   string  `json:"address"`
} 