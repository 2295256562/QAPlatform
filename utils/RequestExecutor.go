package utils

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/oliveagle/jsonpath"
	"strconv"
	"strings"
	"time"
)

type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type ApiCaseStr struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Parameters  string `json:"parameters"` // body参数
	Headers     string `json:"headers"`
	Query       string `json:"query"`
	Asserts     string `json:"asserts"`
	Extract     string `json:"extract"`
	Remark      string `json:"remark"`
	InterfaceId int    `json:"interface_id"`
	EnvId       int    `json:"env_id"`
	EnvName     string `json:"env_name"`

	CreatedBy    int `json:"created_by"`
	ModifiedBy   int `json:"modified_by"`
	CreatedTime  int `json:"created_time"`
	ModifiedTime int `json:"modified_time"`
	ProjectId    int `json:"project_id"`

	InterfaceName string `json:"interface_name"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Domain        string `json:"domain"`

	GVars      string `json:"g_vars"` // 全局变量
	EnvHeaders string `json:"env_headers"`
}

type ApiCase struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Parameters  string    `json:"parameters"` // body参数
	Headers     []Headers `json:"headers"`
	Query       string    `json:"query"`
	Asserts     []Asserts `json:"asserts"`
	Extract     []Extract `json:"extract"`
	Remark      string    `json:"remark"`
	InterfaceId int       `json:"interface_id"`
	EnvId       int       `json:"env_id"`
	EnvName     string    `json:"env_name"`

	CreatedBy    int `json:"created_by"`
	ModifiedBy   int `json:"modified_by"`
	CreatedTime  int `json:"created_time"`
	ModifiedTime int `json:"modified_time"`
	ProjectId    int `json:"project_id"`

	InterfaceName string `json:"interface_name"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Domain        string `json:"domain"`

	GVars      map[string]interface{} `json:"g_vars"` // 全局变量
	EnvHeaders map[string]interface{} `json:"env_headers"`
}

// 测试结果
type ApiCaseResult struct {
	CaseName           string                 `json:"case_name"`
	CaseId             int                    `json:"case_id"`
	InterfaceId        int                    `json:"interface_id"`
	EnvName            string                 `json:"env_name"`
	Method             string                 `json:"method"`
	SuiteId            string                 `json:"suite_id"`
	Url                string                 `json:"url"`
	ResultType         int                    `json:"result_type"`
	RequestHeaders     map[string]interface{} `json:"request_headers"`
	RequestQuery       string                 `json:"request_query"`
	RequestBodyType    string                 `json:"request_body_type"`
	RequestBody        string                 `json:"request_body"`
	ResponseStatusCode int                    `json:"response_status_code"`
	ResponseBody       string                 `json:"response_body"`
	ResponseHeaders    map[string]interface{} `json:"response_headers"`
	ResponseTime       int                    `json:"response_time"`

	ResponseAsserts  []*AssertsResult `json:"response_asserts"`
	ResponseExtracts []ExtractResult  `json:"response_extracts"`
	Exception        string           `json:"exception"`
	ProjectId        int              `json:"project_id"`
	CreatedBy        int              `json:"created_by"`
	Id               int              `json:"id"`
	CreatedTime      int              `json:"created_time"`
}

// 断言结构体
type Asserts struct {
	AssertType string `json:"assert_type"` // 断言类型
	Check      string `json:"check"`       // 断言表达式
	Expect     string `json:"expect"`      // 期望结果
	Comparator string `json:"comparator"`  //期望关系
}

// 断言结果结构体
type AssertsResult struct {
	AssertType string `json:"assert_type"`
	Check      string `json:"check"`      // 断言表达式
	Expect     string `json:"expect"`     // 期望结果
	Comparator string `json:"comparator"` //期望关系
	RealType   string `json:"real_type"`  // 实际值类型
	RealValue  string `json:"real_value"` // 实际值内容
	Result     bool   `json:"result"`     // 结果
}

// 提取参数结构体
type Extract struct {
	ExtractExpress string `json:"extract_express"`
	VarName        string `json:"var_name"`
}

// 提取参数结果结构体
type ExtractResult struct {
	ExtractExpress string `json:"extract_express"`
	VarName        string `json:"var_name"`
	VarValue       string `json:"var_value"`
	VarType        string `json:"var_type"`
}

type CaseLog struct {
	Level       string `json:"level"`
	Msg         string `json:"msg"`
	CreatedTime int    `json:"created_time"`
	ReportId    int    `json:"case_id"`
}

type Cases struct {
	CaseLog []CaseLog
}

// 执行接口自动
func (c *Cases) RequestExecutor(apiCase *ApiCase) (result ApiCaseResult, err error) {
	result.EnvName = apiCase.EnvName
	result.InterfaceId = apiCase.InterfaceId
	result.ProjectId = apiCase.ProjectId
	result.CaseId = apiCase.Id
	result.CaseName = apiCase.Name
	result.Url = apiCase.Url
	result.CreatedBy = apiCase.CreatedBy
	result.CreatedTime = int(time.Now().Unix())
	// 处理headers参数
	headers := c.applyHeaders(apiCase)
	fmt.Println(headers)
	// 处理query参数
	query := c.applyQueryParameters(apiCase)
	fmt.Println(query)
	result.RequestQuery = query
	// 得到请求方法
	method := apiCase.Method
	result.Method = method
	// 处理url
	url := apiCase.Domain + apiCase.Url
	result.Url = url

	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "用例名称：" + result.CaseName, CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "用例ID：" + Strval(result.CaseId), CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "请求地址：" + url, CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "请求方式：" + method, CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "用例环境：" + apiCase.EnvName, CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "当前环境的全局变量：" + MapToJson(apiCase.GVars), CreatedTime: int(time.Now().Unix())})

	// 得到Body参数类型
	bodyType := apiCase.Type
	result.RequestBodyType = bodyType

	var resp *requests.Response
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始发送请求", CreatedTime: int(time.Now().Unix())})

	switch method {
	case "GET":
		resp, err = requests.Get(url + query)
	case "POST":
		parameters := apiCase.Parameters
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始处理body参数", CreatedTime: int(time.Now().Unix())})
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "处理前：" + parameters, CreatedTime: int(time.Now().Unix())})
		body := replaceKeyFromMap(parameters, apiCase.GVars)
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "处理后：" + body, CreatedTime: int(time.Now().Unix())})
		result.RequestBody = body
		if bodyType == "form" {
			result.RequestBodyType = "form"
			mapData, _ := JsonToMap(body)
			resp, err = requests.Post(url, mapData)
		}
		if bodyType == "json" {
			var data map[string]interface{}
			result.RequestBodyType = "json"
			if err1 := json.Unmarshal([]byte(body), &data); err1 == nil {
				resp, err = requests.PostJson(url, data)
			} else {
				c.CaseLog = append(c.CaseLog, CaseLog{Level: "error", Msg: "转换body为json出错", CreatedTime: int(time.Now().Unix())})
				fmt.Println("字符串转json失败：", err1)
			}
		}
	}
	result.ResponseTime = int(time.Now().Unix()) - result.CreatedTime
	result.ResponseBody = resp.Text()
	result.ResponseStatusCode = resp.R.StatusCode

	resultReqHeaders := make(map[string]interface{}, 10)
	for k, v := range resp.R.Request.Header {
		data := v[0]
		resultReqHeaders[k] = data
	}
	result.RequestHeaders = resultReqHeaders
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "请求头：" + MapToJson(resultReqHeaders), CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "响应结果：" + result.ResponseBody, CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "响应状态码：" + Strval(result.ResponseStatusCode), CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "消耗时间为：" + Strval(result.ResponseTime), CreatedTime: int(time.Now().Unix())})

	// 遍历响应头
	resultHeaders := make(map[string]interface{}, 10)
	for k, v := range resp.R.Header {
		var value = v[0]
		resultHeaders[k] = value
	}
	result.ResponseHeaders = resultHeaders
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "响应头：" + MapToJson(resultHeaders), CreatedTime: int(time.Now().Unix())})
	assertResult := c.handleAssert(apiCase, &result)
	if assertResult {
		result.ResultType = 1
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "测试结果：成功", CreatedTime: int(time.Now().Unix())})
	} else {
		result.ResultType = 0
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "error", Msg: "测试结果：失败", CreatedTime: int(time.Now().Unix())})
	}

	c.handleExtract(apiCase, &result)
	return
}

// 处理headers函数
func (c *Cases) applyHeaders(apiCase *ApiCase) (headerStr string) {
	// TODO 合并两个headers
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始处理请求headers", CreatedTime: int(time.Now().Unix())})
	headers := apiCase.Headers
	headersTemp := make(map[string]interface{})
	if headers != nil {
		for _, header := range headers {
			if header.Key != "" || header.Value != "" {
				value := replaceKeyFromMap(header.Value, apiCase.GVars)
				headersTemp[header.Key] = value
			}
		}
	}
	headerStr = MapToJson(headersTemp)
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "处理后的headers:" + headerStr, CreatedTime: int(time.Now().Unix())})
	return
}

// 处理query参数
func (c *Cases) applyQueryParameters(apiCase *ApiCase) (query string) {
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始处理请求query请求参数", CreatedTime: int(time.Now().Unix())})
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "query参数：" + apiCase.Query, CreatedTime: int(time.Now().Unix())})
	query = replaceKeyFromMap(apiCase.Query, apiCase.GVars)
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "处理后的query参数：" + query, CreatedTime: int(time.Now().Unix())})
	return
}

// 断言函数
func (c *Cases) handleAssert(apiCase *ApiCase, result *ApiCaseResult) (flag bool) {
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始进行断言", CreatedTime: int(time.Now().Unix())})
	flag = true
	asserts := apiCase.Asserts
	var assertsList = make([]*AssertsResult, 0)
	for _, assert := range asserts {
		str, _ := json.Marshal(assert)
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "进行：" + Strval(str) + "断言", CreatedTime: int(time.Now().Unix())})
		assertResult := c.Assertion(&assert, result)
		assertsList = append(assertsList, &assertResult)
		if !assertResult.Result {
			flag = false
		}
	}
	result.ResponseAsserts = assertsList
	return
}

func (c *Cases) Assertion(assertion *Asserts, result *ApiCaseResult) (assertResult AssertsResult) {
	assertResult.AssertType = assertion.AssertType
	assertResult.Check = assertion.Check
	assertResult.Expect = assertion.Expect
	assertResult.Comparator = assertion.Comparator

	switch assertion.AssertType {
	case "response_json":
		extract, err := JsonPathExtract(result.ResponseBody, assertion.Check)
		if err != nil {
			assertResult.RealType = "null"
			assertResult.RealValue = ""
		}
		realType := getObjRealType(extract)
		if realType == "null" {
			assertResult.RealType = realType
			assertResult.RealValue = ""
		} else {
			assertResult.RealType = realType
			assertResult.RealValue = Strval(extract)
		}
		break
	case "status_code":
		realType := getObjRealType(result.ResponseStatusCode)
		if realType == "null" {
			assertResult.RealType = "null"
			assertResult.RealValue = ""
		} else {
			assertResult.RealType = "number"
			assertResult.RealValue = Strval(result.ResponseStatusCode)
		}
		break
	case "response":
		if result.ResponseBody == "" {
			assertResult.RealType = "null"
		} else {
			assertResult.RealType = "string"
		}
		assertResult.RealValue = result.ResponseBody
	}

	flag, err := getAssertionResult(assertion, assertResult.RealType, assertResult.RealValue)

	if err != nil {
		assertResult.Result = false
	}
	assertResult.Result = flag
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "断言结果：" + Strval(assertResult.Result), CreatedTime: int(time.Now().Unix())})

	return
}

// 参数提取
func (c *Cases) handleExtract(apiCase *ApiCase, result *ApiCaseResult) {
	c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "开始参数提取", CreatedTime: int(time.Now().Unix())})

	var extractList = make([]ExtractResult, 0)
	extract := apiCase.Extract
	for _, e := range extract {
		str, _ := json.Marshal(e)
		c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "进行参数提取" + Strval(str), CreatedTime: int(time.Now().Unix())})
		value, err := JsonPathExtract(result.ResponseBody, e.ExtractExpress)
		if err != nil {
			c.CaseLog = append(c.CaseLog, CaseLog{Level: "error", Msg: "提取参数失败", CreatedTime: int(time.Now().Unix())})
			extractList = append(extractList, ExtractResult{VarName: e.VarName, ExtractExpress: e.ExtractExpress, VarValue: "null", VarType: "null"})
		} else {
			c.CaseLog = append(c.CaseLog, CaseLog{Level: "info", Msg: "提取参数成功，提取的值：" + Strval(value), CreatedTime: int(time.Now().Unix())})
			extractList = append(extractList, ExtractResult{VarName: e.VarName, ExtractExpress: e.ExtractExpress, VarValue: Strval(value), VarType: getObjRealType(value)})
		}
	}
	result.ResponseExtracts = extractList
}

// 获取断言结果
func getAssertionResult(ass *Asserts, realType, realValue string) (flag bool, err error) {
	comparator := ass.Comparator
	switch comparator {
	case "相等":
		if ass.Expect == realValue {
			return true, nil
		}
	case "大于":
		if realType == "number" {
			var realva, expect int
			realva, err = strconv.Atoi(realValue)
			expect, err = strconv.Atoi(ass.Expect)
			if err != nil {
				return false, err
			}
			if realva > expect {
				return true, nil
			} else {
				return false, nil
			}
		}
	case "包含":
		if strings.Contains(realValue, ass.Expect) {
			return true, nil
		}
	}
	return
}

// jsonpath表达式提取器
func JsonPathExtract(response, extractExpression string) (res interface{}, err error) {
	var jsonData interface{}
	json.Unmarshal([]byte(response), &jsonData)

	res, err = jsonpath.JsonPathLookup(jsonData, extractExpression)
	if err != nil {
		return
	}
	return
}

// 字符串正则替换变量
func replaceKeyFromMap(str string, vars map[string]interface{}) (result string) {
	for k, v := range vars {
		str = strings.ReplaceAll(str, `${`+k+`}`, Strval(v))
	}
	return str
}

// 获取参数类型
func getObjRealType(obj interface{}) (kind string) {
	if obj == nil {
		return "null"
	}
	switch obj.(type) { //多选语句switch
	case string:
		return "string"
	case int:
		return "number"
	case bool:
		return "boolean"
	case float64:
		return "number"
	default:
		return "string"
	}
	return
}

func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

func JsonToMap(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}

	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}

	return m, nil
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}
func intToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
func boolToString(b bool) string {
	if b {
		return "true"
	} else {
		return "false"
	}
}

func toString(arg interface{}) string {
	switch arg.(type) {
	case bool:
		return boolToString(arg.(bool))
	case float32:
		return floatToString(float64(arg.(float32)))
	case float64:
		return floatToString(arg.(float64))
	case int:
		return intToString(int64(arg.(int)))
	case int8:
		return intToString(int64(arg.(int8)))
	case int16:
		return intToString(int64(arg.(int16)))
	case int32:
		return intToString(int64(arg.(int32)))
	case int64:
		return intToString(int64(arg.(int64)))
	default:
		return fmt.Sprint(arg)
	}
}
