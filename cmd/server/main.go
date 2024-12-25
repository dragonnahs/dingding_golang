package main

import (
    "flag"
    "log"
    
    "your-project/pkg/config"
    "your-project/internal/router"
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