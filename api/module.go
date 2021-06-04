package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddModule(c *gin.Context) {
	var moduleStr *model.Module
	err := c.ShouldBindJSON(&moduleStr)
	userId := c.MustGet("id").(int)

	moduleStr.CreatedBy = userId
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("输入参数异常")))
		return
	}

	exist := model.CheckModuleNameExist(moduleStr.Name, moduleStr.ProjectId)
	if exist {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("同项目内模块名称不可重复")))
		return
	}

	flag := model.AddModule(moduleStr)
	if flag == false {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("创建模块失败")))
		return
	}
	utils.ResponseSuccess(c, "创建成功")
	return

}

func ModelList(c *gin.Context) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()
	name := c.Query("name")
	projectId := c.Query("project_id")

	maps := make(map[string]interface{})

	if projectId == "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}
	maps["state"] = 1
	maps["project_id"] = projectId

	if name != "" {
		maps["name"] = name
	}

	list, count := model.GetModuleList(pageSize, page, maps)
	data := make(map[string]interface{})
	data["rows"] = list
	data["count"] = count
	utils.ResponseSuccess(c, data)
}

func Modules(c *gin.Context) {
	projectId := com.StrTo(c.Query("project_id")).MustInt()
	if projectId == 0 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}
	modules, err := model.GetModules(projectId)
	if err != nil {
		utils.ResponseError(c, 500, err)
		return
	}
	utils.ResponseSuccess(c, modules)
	return
}

func DelModule(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	flag := model.ModuleDel(id)
	if flag == false {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("删除失败")))
		return
	}
	utils.ResponseSuccess(c, "删除成功")
	return
}

func ModuleEdit(c *gin.Context) {
	var data *model.Module

	err := c.ShouldBindJSON(&data)
	userId := c.MustGet("id").(int)
	data.ModifiedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数异常")))
		return
	}

	flag := model.ModuleEdit(data)
	if flag == false {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("修改模块信息失败")))
		return
	}
	utils.ResponseSuccess(c, "修改成功")
	return
}
