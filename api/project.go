package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

}
