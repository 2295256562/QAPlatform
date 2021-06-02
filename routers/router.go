package routers

import (
	"QAPlatform/api"
	"QAPlatform/middleware"
	"QAPlatform/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	V1 := r.Group("/api/v1")
	V1.Use(middleware.RequestLog())
	{
		V1.POST("/login", api.Login)
		V1.Use(middleware.JwtToken())
	}
	r.Run(utils.HttpPort)
}
