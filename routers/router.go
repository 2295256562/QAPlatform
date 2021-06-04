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
		// 用户相关接口
		V1.POST("/login", api.Login)
		V1.POST("/register", api.Register)
		V1.Use(middleware.JWTAuthMiddleware())
		V1.GET("/users", api.QueryUserListOnRole)

		// 项目相关接口
		V1.POST("/add_project", api.CreateProject)
		V1.GET("/project_list", api.ProjectList)
		V1.GET("/project_detail", api.ProjectDetail)
		V1.PUT("/project/:id", api.ProjectEdit)
		V1.DELETE("/project/:id", api.ProjectDel)

		// 模块接口
		V1.GET("/module_list", api.ModelList)
		V1.POST("/module_add", api.AddModule)
		V1.GET("/modules", api.Modules)
		V1.DELETE("/module/:id", api.DelModule)
		V1.POST("/module_update", api.ModuleEdit)

		// 接口管理接口
		V1.POST("/inter_add", api.AddInter)
	}
	r.Run(utils.HttpPort)
}
