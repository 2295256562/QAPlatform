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
	r.Use(middleware.Cors())
	{
		V1.POST("/login", api.Login)
		V1.POST("/register", api.Register)
		V1.Use(middleware.JWTAuthMiddleware())
		V1.POST("/add_project", api.CreateProject)
	}
	r.Run(utils.HttpPort)
}
