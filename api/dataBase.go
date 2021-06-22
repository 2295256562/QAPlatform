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
