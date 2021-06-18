package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"encoding/json"
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

func InterfaceCaseDebug(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()
	userId := c.MustGet("id").(int)

	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用例id不能为空")))
		return
	}
	info, err := model.CaseInfo(Id)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询用例详情失败")))
		return
	}

	var header []utils.Headers
	var asserts []utils.Asserts
	var extract []utils.Extract
	var gVars map[string]interface{}
	var env_headers map[string]interface{}

	fmt.Println(info.Extract)

	if info.Headers != "{}" {
		multiErr := json.Unmarshal([]byte(info.Headers), &header)
		if multiErr != nil {
			utils.ResponseError(c, 500, errors.New(fmt.Sprint("转换headers出错")))
			return
		}
	}

	if info.Asserts != "[]" {
		multiErr1 := json.Unmarshal([]byte(info.Asserts), &asserts)
		if multiErr1 != nil {
			utils.ResponseError(c, 500, errors.New(fmt.Sprint("转换Asserts出错")))
			return
		}
	}

	multiErr2 := json.Unmarshal([]byte(info.Extract), &extract)
	if multiErr2 != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("转换Extract出错")))
		return
	}

	multiErr3 := json.Unmarshal([]byte(info.GVars), &gVars)
	if multiErr3 != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("转换全局变量出错")))
		return
	}

	multiErr4 := json.Unmarshal([]byte(info.EnvHeaders), &env_headers)
	if multiErr4 != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("转换全局请求头出错")))
		return
	}
	caseExecution := &utils.Cases{}
	result, err := caseExecution.RequestExecutor(&utils.ApiCase{
		Id:            info.Id,
		Name:          info.Name,
		Type:          info.Type,
		Parameters:    info.Parameters,
		Headers:       header,
		Query:         info.Query,
		Asserts:       asserts,
		Extract:       extract,
		Remark:        info.Remark,
		InterfaceId:   info.InterfaceId,
		EnvId:         info.EnvId,
		EnvName:       info.EnvName,
		CreatedBy:     info.CreatedBy,
		ModifiedBy:    info.ModifiedBy,
		CreatedTime:   info.CreatedTime,
		ModifiedTime:  info.ModifiedTime,
		ProjectId:     info.ProjectId,
		InterfaceName: info.InterfaceName,
		Url:           info.Url,
		Method:        info.Method,
		Domain:        info.Domain,
		GVars:         gVars,
		EnvHeaders:    env_headers,
	})

	// 存储测试结果
	ass, err := json.Marshal(result.ResponseAsserts)
	ext, err := json.Marshal(result.ResponseExtracts)
	res, err1 := model.AddCaseResult(&model.ApiCaseResultStr{
		CaseName:           result.CaseName,
		CaseId:             result.CaseId,
		InterfaceId:        result.InterfaceId,
		EnvName:            result.EnvName,
		SuiteId:            0,
		Method:             result.Method,
		Url:                result.Url,
		ResultType:         result.ResultType,
		RequestHeaders:     utils.MapToJson(result.RequestHeaders),
		RequestQuery:       result.RequestQuery,
		RequestBodyType:    result.RequestBodyType,
		RequestBody:        result.RequestBody,
		ResponseStatusCode: result.ResponseStatusCode,
		ResponseBody:       result.ResponseBody,
		ResponseHeaders:    utils.MapToJson(result.ResponseHeaders),
		ResponseTime:       result.ResponseTime,
		ResponseAsserts:    utils.Strval(ass),
		ResponseExtracts:   utils.Strval(ext),
		Exception:          result.Exception,
		ProjectId:          result.ProjectId,
		CreatedBy:          userId,
	})

	if err1 != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("存储执行结果失败")))
		return
	}
	// 存储执行日志
	for _, log := range caseExecution.CaseLog {
		model.AddCaseLog(&model.CaseLog{Msg: log.Msg, Level: log.Level, ReportId: res})
	}

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("执行用例失败")))
		return
	}
	utils.ResponseSuccess(c, res)
	return
}

func InterfaceCaseResult(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()
	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("测试报告id不能为空")))
		return
	}
	result, err := model.QueryCaseResult(Id)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询测试报告出错")))
		return
	}
	utils.ResponseSuccess(c, result)
	return
}

func InterfaceCaseLog(c *gin.Context) {
	Id := com.StrTo(c.Query("id")).MustInt()

	if Id < 1 {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("报告id不能为空")))
		return
	}

	logs := model.QueryCaseLogByReportId(Id)
	utils.ResponseSuccess(c, logs)
	return
}
