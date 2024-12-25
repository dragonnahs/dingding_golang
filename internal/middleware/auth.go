/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 14:21:29
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 14:21:39
 * @FilePath: \dingding_golang\internal\middleware\auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
    "net/http"
    "dingding_golang/pkg/utils"

    "github.com/gin-gonic/gin"

)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": 401,
                "message": "未授权访问",
            })
            c.Abort()
            return
        }

        // 验证 token
        claims, err := utils.ParseToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": 401,
                "message": "无效的 token",
            })
            c.Abort()
            return
        }

        // 将用户信息存储到上下文
        c.Set("userId", claims.UserId)
        c.Next()
    }
} 