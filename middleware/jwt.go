package middleware

import (
	"github.com/gin-gonic/gin"
	"memorial01/pkg/utils"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200 //状态码

		token := c.GetHeader("Authorization") //获得token
		if token == "" {
			code = 404 //没有取得token
		} else {
			claim, err := utils.ParseToken(token) //验证token是否正确
			if err != nil {
				code = 403 //说明这个token是无权限的,是假的
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //这个token已经过期了
			}
		}
		if code != 200 {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    "Token解析错误",
			})
			c.Abort() //终止
			return
		}
		c.Next() //挂起

	}
}
