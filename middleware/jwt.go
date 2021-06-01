package middleware

import (
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// jwt鉴权中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		if tokenHerder == "" {
			utils.ResponseError(c, 401, errors.New(fmt.Sprintf("请登录")))
		}
	}
}
