/*
 * @Author: shanlonglong danlonglong@weimiao.cn
 * @Date: 2024-12-25 14:25:00
 * @LastEditors: shanlonglong danlonglong@weimiao.cn
 * @LastEditTime: 2024-12-25 14:23:33
 * @FilePath: \dingding_golang\pkg\utils\token.go
 * @Description: Token 处理工具
 */
package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserId   string `json:"userId"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

var jwtSecret = []byte("your_jwt_secret") // 在实际应用中应该从配置中读取

func GenerateToken(userId, username string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(24 * time.Hour)

    claims := Claims{
        UserId:   userId,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expireTime),
            IssuedAt:  jwt.NewNumericDate(nowTime),
            Subject:   userId,
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
} 