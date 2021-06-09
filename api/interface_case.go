package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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

	if data.Name == "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用例名称不能为空")))
		return
	}

	if data.EnvId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境id不能为空")))
		return
	}

	err = model.InterfaceCaseAdd(data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("添加接口失败")))
		return
	}

	utils.ResponseSuccess(c, "添加成功")
	return
}

func InterfaceCaseList(c *gin.Context) {
	var data *model.InterfaceQueryDto

	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数异常")))
		return
	}
	list := model.CaseList(data)
	resp := make(map[string]interface{})
	resp["rows"] = list
	resp["count"] = len(list)
	utils.ResponseSuccess(c, resp)
	return
}

func InterfaceDetail(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()

	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用例id不能为空")))
		return
	}

	detail, err := model.CaseDetail(Id)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询用例详情失败")))
		return
	}
	utils.ResponseSuccess(c, detail)
	return
}

func InterfaceCaseEdit(c *gin.Context) {
	var data *model.InterfaceCase
	err := c.ShouldBindJSON(&data)
	userId := c.MustGet("id").(int)
	data.ModifiedBy = userId

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

	if data.Name == "" {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用例名称不能为空")))
		return
	}

	if data.EnvId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("环境id不能为空")))
		return
	}

	err = model.CaseEdit(data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("修改接口失败")))
		return
	}

	utils.ResponseSuccess(c, "修改成功")
	return
}
