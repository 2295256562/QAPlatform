package api

import "C"
import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
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
	flag := model.CheckInterfaceNameExists(data.Name, data.ProjectId)
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

func ListInterByModuleId(c *gin.Context) {
	moduleId := com.StrTo(c.Query("module_id")).MustInt()
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()

	if moduleId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("模块id不可为空")))
		return
	}

	list, count, err := model.FindListByModuleId(pageSize, page, moduleId)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询列表出错")))
		return
	}
	data := make(map[string]interface{})
	data["rows"] = list
	data["count"] = count
	utils.ResponseSuccess(c, data)
	return
}

func InterfaceList(c *gin.Context) {
	projectId := com.StrTo(c.Query("project_id")).MustInt()
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()

	if projectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}

	list, count, err := model.InterList(pageSize, page, projectId)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询列表出错")))
		return
	}
	data := make(map[string]interface{})
	data["rows"] = list
	data["count"] = count
	utils.ResponseSuccess(c, data)
	return
}

func InterDetail(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()

	if id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("接口id不可为空")))
		return
	}
	detail, err := model.InterDetail(id)

	if detail.Id == 0 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("接口不存在")))
		return
	}

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("接口详情出错")))
		return
	}
	utils.ResponseSuccess(c, detail)
	return
}

func InterDel(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	fmt.Println(id)
	// TODO 查询当前接口下是否存在用例，如果有用例不可删除
}

func InterEdit(c *gin.Context) {
	var data *model.InterfaceAdd
	err := c.ShouldBindJSON(&data)

	userId := c.MustGet("id").(int)
	data.ModifiedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数有误")))
		return
	}
	err = model.InterUpdate(data)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("修改接口失败")))
		return
	}
	utils.ResponseSuccess(c, "修改接口成功")
	return
}
