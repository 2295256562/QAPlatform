package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"github.com/unknwon/com"
	"os"
	"path"
	"strconv"
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
		model.AddCaseLog(&model.CaseLog{Msg: log.Msg, Level: log.Level, CreatedTime: log.CreatedTime, ReportId: res})
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

// InterfaceExport 导出测试用例
func InterfaceExport(c *gin.Context) {
	var data *model.InterfaceQueryDto

	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("参数异常")))
		return
	}
	list := model.CaseExport(data)

	titleList := []string{"ID", "用例名称", "接口名称", "环境名称", "请求路径", "请求方式",
		"请求Header", "请求query", "请求body", "body类型", "断言信息", "提取参数", "备注"}

	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		//居中显示
		cell.GetStyle().Alignment.Horizontal = "center"
		cell.GetStyle().Alignment.Vertical = "center"
	}
	// 插入内容
	for _, v := range list {
		row := sheet.AddRow()
		row.AddCell().SetInt(v.Id)
		row.AddCell().SetString(v.Name)
		row.AddCell().SetString(v.InterfaceName)
		row.AddCell().SetString(v.EnvName)
		row.AddCell().SetString(v.Domain + v.Url)
		row.AddCell().SetString(v.Method)
		row.AddCell().SetString(v.Headers)
		row.AddCell().SetString(v.Query)
		row.AddCell().SetString(v.Parameters)
		row.AddCell().SetString(v.Type)
		row.AddCell().SetString(v.Asserts)
		row.AddCell().SetString(v.Extract)
		row.AddCell().SetString(v.Remark)
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=%s.xlsx", "cases")
	c.Writer.Header().Set("Content-Disposition", disposition)
	_ = file.Write(c.Writer)
}

func InterfaceImport(c *gin.Context) {

	var createdBy int
	var ProjectId int

	file, _ := c.FormFile("file")
	dst := path.Join("./upload", file.Filename)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	}
	xlsx, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 获取excel中具体的列的值
	rows, _ := xlsx.GetRows("Sheet" + "1")
	for key, row := range rows {
		if key == 0 {
			continue
		}
		InterfaceId, err := strconv.Atoi(row[8])
		if err != nil {
			utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
			return
		}
		EnvId, err := strconv.Atoi(row[9])
		if err != nil {
			utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
			return
		}
		// 去掉标题行
		if key > 0 {
			cases := model.InterfaceCase{Name: row[0], Type: row[1], Parameters: row[2], Headers: row[3], Query: row[4],
				Asserts: row[5], Extract: row[6], Remark: row[7], InterfaceId: InterfaceId, EnvId: EnvId, CreatedBy: createdBy, ProjectId: ProjectId}
			err := model.InterfaceCaseAdd(&cases)
			if err != nil {
				utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
				return
			}
		}
	}
	utils.ResponseSuccess(c, "导入成功")
	return
}

func InterfaceDownloadTemplate(c *gin.Context) {
	titleList := []string{"用例名称", "body类型", "body参数", "请求header", "请求query",
		"断言信息", "参数提取", "备注", "接口id", "环境id", "创建人员", "项目id"}

	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet("Sheet1")
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
		//居中显示
		cell.GetStyle().Alignment.Horizontal = "center"
		cell.GetStyle().Alignment.Vertical = "center"
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename=%s.xlsx", "cases")
	c.Writer.Header().Set("Content-Disposition", disposition)
	_ = file.Write(c.Writer)
}
