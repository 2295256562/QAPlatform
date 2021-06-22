package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func AddDataBase(c *gin.Context) {
	var data *model.DataBase
	err := c.ShouldBindJSON(&data)

	userId := c.MustGet("id").(int)
	data.CreatedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数有误")))
		return
	}
	err = model.AddDataBase(data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	}
	utils.ResponseSuccess(c, "新增成功")
	return
}

func DataBaseEdit(c *gin.Context) {
	var data *model.DataBase
	err := c.ShouldBindJSON(&data)

	userId := c.MustGet("id").(int)
	data.ModifiedBy = userId

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数有误")))
		return
	}
	err = model.EditDataBase(data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	}
	utils.ResponseSuccess(c, "修改成功")
	return
}

func DataBaseDetail(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()
	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("id不能为空")))
		return
	}
	detail, err := model.DetailDataBase(Id)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	}
	utils.ResponseSuccess(c, detail)
	return
}

func DataBaseList(c *gin.Context) {
	projectId := com.StrTo(c.Query("project_id")).MustInt()
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "10")).MustInt()

	if projectId < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("项目id不可为空")))
		return
	}
	lists, count, err := model.DataBaseLists(pageSize, page, projectId)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询列表出错")))
		return
	}
	data := make(map[string]interface{})
	data["rows"] = lists
	data["count"] = count
	utils.ResponseSuccess(c, data)
	return
}

func DataBaseDel(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	if id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("id不能为空")))
		return
	}
	err := model.DataBaseDelete(id)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	}
	utils.ResponseSuccess(c, "删除成功")
	return
}
