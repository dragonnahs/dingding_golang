package controller

import (
	"dingding_golang/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
    return &AuthController{}
}

func (c *AuthController) GenerateToken(ctx *gin.Context) {
    var req struct {
        UserId   string `json:"userId" binding:"required"`
        Username string `json:"username" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "code": 400,
            "message": "无效的请求参数",
        })
        return
    }

    token, err := utils.GenerateToken(req.UserId, req.Username)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "code": 500,
            "message": "生成token失败",
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "code": 200,
        "message": "success",
        "data": gin.H{
            "token": token,
        },
    })
} 