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
	r.Use(middleware.Cors())
	V1 := r.Group("/api/v1")
	V1.Use(middleware.RequestLog())
	{
		// 用户相关接口
		V1.POST("/login", api.Login)
		V1.POST("/register", api.Register)
		V1.Use(middleware.JWTAuthMiddleware())
		V1.GET("/users", api.QueryUserListOnRole)
		V1.GET("/user_all", api.UserAll)

		// 项目相关接口
		V1.POST("/add_project", api.CreateProject)
		V1.GET("/project_list", api.ProjectList)
		V1.GET("/project_detail", api.ProjectDetail)
		V1.PUT("/project/:id", api.ProjectEdit)
		V1.DELETE("/project/:id", api.ProjectDel)
		V1.GET("/projects", api.Projects)

		// 模块接口
		V1.GET("/module_list", api.ModelList)
		V1.POST("/module_add", api.AddModule)
		V1.GET("/modules", api.Modules)
		V1.DELETE("/module/:id", api.DelModule)
		V1.POST("/module_update", api.ModuleEdit)

		// 环境管理接口
		V1.POST("/env_add", api.EnvAdd)
		V1.GET("/env_list", api.EnvList)
		V1.GET("/envs", api.Envs)
		V1.DELETE("/env_del/:id", api.EnvDel)
		V1.GET("/env_info", api.EnvDetail)
		V1.POST("/env_edit", api.EnvEdit)

		// 接口管理接口
		V1.POST("/inter_add", api.AddInter)
		V1.GET("/inter_list", api.ListInterByModuleId)
		V1.GET("/inters", api.InterfaceList)
		V1.GET("/inter_all", api.Inters)
		V1.GET("/inter", api.InterDetail)
		V1.POST("/inter_edit", api.InterEdit)
		V1.DELETE("/inter_del/:id", api.InterDel)

		// 用例相关接口
		V1.POST("/add_case", api.AddCase)
		V1.GET("/case_list", api.InterfaceCaseList)
		V1.GET("/case_detail", api.InterfaceDetail)
		V1.POST("/case_edit", api.InterfaceCaseEdit)
		V1.GET("/case_debug", api.InterfaceCaseDebug)
		V1.GET("/case_log", api.InterfaceCaseLog)
		V1.GET("/case_result", api.InterfaceCaseResult)
		V1.GET("/case_export", api.InterfaceExport)
	}
	r.Run(utils.HttpPort)
}
