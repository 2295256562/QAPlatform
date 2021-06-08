package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func EnvList(c *gin.Context) {
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()
	projectId := com.StrTo(c.Query("project_id")).MustInt()

	if projectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}

	list, count, err := model.EnvironmentLists(pageSize, page, projectId)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询出错")))
		return
	}

	data := make(map[string]interface{})
	data["rows"] = list
	data["count"] = count
	utils.ResponseSuccess(c, data)
	return
}

func Envs(c *gin.Context) {
	projectId := com.StrTo(c.Query("project_id")).MustInt()

	if projectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}
	environments, err := model.Environments(projectId)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}
	utils.ResponseSuccess(c, environments)
	return
}

func EnvAdd(c *gin.Context) {
	var AddEnv *model.Environment
	err := c.ShouldBindJSON(&AddEnv)

	userId := c.MustGet("id").(int)
	AddEnv.CreatedBy = userId

	if AddEnv.ProjectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不能为空")))
		return
	}

	if AddEnv.Name == "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境名称不能为空")))
		return
	}

	if AddEnv.Domain == "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境域名不能为空")))
		return
	}

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("输入参数有误")))
		return
	}
	exist := model.CheckNameExist(AddEnv.Name, AddEnv.ProjectId)
	if exist {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境名称不可重复，请更换名称")))
		return
	}

	environment, err := model.AddEnvironment(AddEnv)
	if err != nil || !environment {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("创建环境失败")))
		return
	}
	utils.ResponseSuccess(c, "创建成功")
	return
}

func EnvDel(c *gin.Context) {
	Id := com.StrTo(c.Param("id")).MustInt()

	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境id不能为空")))
		return
	}

	flag := model.EnvironmentDel(Id)
	if !flag {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("删除失败")))
		return
	}
	utils.ResponseSuccess(c, "删除成功")
	return
}

func EnvDetail(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()

	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境id不能为空")))
		return
	}

	detail, err := model.EnvironmentDetail(Id)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("获取项目详情报错")))
		return
	}
	utils.ResponseSuccess(c, detail)
	return
}

func EnvEdit(c *gin.Context) {
	var env *model.Environment
	err := c.ShouldBindJSON(&env)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("绑定编辑环境参数失败")))
		return
	}
	edit, err := model.EnvironmentEdit(env)
	if err != nil || !edit {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("编辑项目失败")))
		return
	}
	utils.ResponseSuccess(c, "修改成功")
	return
}
