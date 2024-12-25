/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 13:53:37
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 14:14:35
 * @FilePath: \dingding_golang\cmd\server\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
    "flag"
    "log"
    "dingding_golang/pkg/config"
    "dingding_golang/internal/router"

)

func main() {
    var env string
    flag.StringVar(&env, "env", "dev", "运行环境 (dev/prod)")
    flag.Parse()

    // 加载配置
    if err := config.Load(env); err != nil {
        log.Fatalf("加载配置失败: %v", err)
    }

    // 初始化路由
    r := router.InitRouter()
    
    // 启动服务
    if err := r.Run(config.Get().Server.Address); err != nil {
        log.Fatalf("启动服务失败: %v", err)
    }
} 