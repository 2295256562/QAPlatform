package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// CreateProject 创建项目
func CreateProject(c *gin.Context) {
	var AddProject *model.AddProject
	err := c.BindJSON(&AddProject)

	userId := c.MustGet("id").(int)

	AddProject.CreatedBy = userId
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("输入参数有误")))
		return
	}
	exist := model.CheckProjectExist(AddProject.Name)
	if exist {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目名称重复，请更换名称")))
		return
	}
	flag := model.CreateProject(AddProject)
	if !flag {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("创建项目失败")))
		return
	}
	utils.ResponseSuccess(c, "创建成功")
	return
}

// ProjectList 获取项目列表带分页
func ProjectList(c *gin.Context) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()
	name := c.Query("name")
	maps := make(map[string]interface{})

	maps["state"] = 1
	if name != "" {
		maps["name"] = name
	}

	list, count := model.GetProjectList(pageSize, page, maps)
	data := make(map[string]interface{})
	data["rows"] = list
	data["count"] = count
	utils.ResponseSuccess(c, data)
	return
}

// 获取项目列表不带分页
func Projects(c *gin.Context) {
	projects, err := model.GetProjects()
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数异常")))
		return
	}
	utils.ResponseSuccess(c, projects)
	return
}

func ProjectDetail(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	if id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("请传入项目id")))
		return
	}
	detail := model.GetProjectDetail(id)
	utils.ResponseSuccess(c, detail)
}

func ProjectEdit(c *gin.Context) {
	var data *model.AddProject

	id := com.StrTo(c.Param("id")).MustInt()
	err := c.ShouldBindJSON(&data)
	userId := c.MustGet("id").(int)
	data.ModifiedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数异常")))
		return
	}

	exist := model.CheckProjectExist(data.Name)
	if exist {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目名称重复，请更换名称")))
		return
	}

	flag := model.EditProject(id, data)
	if flag == false {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("修改失败")))
		return
	}
	utils.ResponseSuccess(c, "修改成功")
	return
}

func ProjectDel(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	flag := model.DelProject(id)
	if flag == false {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("删除失败")))
		return
	}
	utils.ResponseSuccess(c, "删除成功")
	return
}
