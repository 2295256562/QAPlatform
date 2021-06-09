package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddCase(c *gin.Context) {
	var data *model.InterfaceCase
	err := c.ShouldBindJSON(&data)
	userId := c.MustGet("id").(int)
	data.CreatedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数有误")))
		return
	}

	if data.InterfaceId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("接口id不能为空")))
		return
	}

	if data.ProjectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不能为空")))
		return
	}

	if data.Name != "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用例名称不能为空")))
		return
	}

	if data.EnvId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境id不能为空")))
		return
	}

}
