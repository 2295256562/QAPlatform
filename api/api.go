package api

import "C"
import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddInter(c *gin.Context) {
	var data *model.InterfaceAdd
	err := c.ShouldBindJSON(&data)

	userId := c.MustGet("id").(int)
	data.CreatedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数有误")))
		return
	}
	flag, err := model.CheckInterfaceNameExists(data.Name, data.ProjectId)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("校验接口名称是否存在报错")))
		return
	}
	if flag {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("同项目类接口名称不可重复")))
		return
	}

	err = model.AddApi(data)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("添加接口失败")))
		return
	}
	utils.ResponseSuccess(c, "新增接口成功")
	return
}
