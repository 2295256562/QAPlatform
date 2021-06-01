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
	r.Use(middleware.JwtToken())
	V1 := r.Group("/api/v1")
	{
		V1.POST("/login", api.Login)
		V1.Use(middleware.RequestLog())
	}
	r.Run(utils.HttpPort)
}
