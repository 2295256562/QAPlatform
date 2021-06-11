package utils

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/oliveagle/jsonpath"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var log *zap.Logger

type Headers struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	CaseName           string `json:"case_name"`
	CaseId             int    `json:"case_id"`
	InterfaceId        int    `json:"interface_id"`
	Method             string `json:"method"`
	SuiteId            string `json:"suite_id"`
	Url                string `json:"url"`
	ResultType         int    `json:"result_type"`
	RequestHeaders     string `json:"request_headers"`
	RequestQuery       string `json:"request_query"`
	RequestBodyType    string `json:"request_body_type"`
	RequestBody        string `json:"request_body"`
	ResponseStatusCode int    `json:"response_status_code"`
	ResponseBody       string `json:"response_body"`
	ResponseHeaders    string `json:"response_headers"`
	ResponseTime       string `json:"response_time"`

	ResponseAsserts  []AssertsResult `json:"response_asserts"`
	ResponseExtracts string          `json:"response_extracts"`
	Exception        string          `json:"exception"`
	ProjectId        int             `json:"project_id"`
	CreatedBy        int             `json:"created_by"`
	ModifiedBy       int             `json:"modified_by"`
	Id               int             `json:"id"`
	CreatedTime      int             `json:"created_time"`
	ModifiedTime     int             `json:"modified_time"`
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

// 执行接口自动
func RequestExecutor(apiCase *ApiCase) (result ApiCaseResult, err error) {
	result.InterfaceId = apiCase.InterfaceId
	result.ProjectId = apiCase.ProjectId
	result.CaseId = apiCase.Id
	result.CaseName = apiCase.Name
	result.CreatedTime = int(time.Now().Unix())

	// 处理headers参数
	headers := applyHeaders(apiCase)
	result.RequestHeaders = headers
	// 处理query参数
	query := applyQueryParameters(apiCase)
	result.RequestQuery = query
	// 得到请求方法
	method := apiCase.Method
	result.Method = method
	// 处理url
	url := apiCase.Domain + apiCase.Url
	fmt.Println("请求地址为：", url)
	// 得到Body参数类型
	bodyType := apiCase.Type
	result.RequestBodyType = bodyType

	var resp *requests.Response

	switch method {
	case "GET":
		resp, err = requests.Get(url + query)
	case "POST":
		parameters := apiCase.Parameters
		body := replaceKeyFromMap(parameters, apiCase.GVars)
		result.RequestBody = body
		fmt.Println("body信息 ===>", body)
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
				fmt.Println("字符串转json失败：", err1)
			}
		}
	}
	result.ResponseBody = resp.Text()
	result.ResponseStatusCode = resp.R.StatusCode
	fmt.Println(resp.Text())
	assertResult := handleAssert(apiCase, &result)
	if assertResult {
		result.ResultType = 1
	} else {
		result.ResultType = 0
	}
	fmt.Println("断言结果：", result.ResponseAsserts)
	fmt.Println("测试结果：", result.ResultType)
	return
}

// 处理headers函数
func applyHeaders(apiCase *ApiCase) (headerStr string) {
	fmt.Println("开始处理headers")
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
	return MapToJson(headersTemp)
}

// 处理query参数
func applyQueryParameters(apiCase *ApiCase) (query string) {
	query = replaceKeyFromMap(apiCase.Query, apiCase.GVars)
	return
}

// 断言函数
func handleAssert(apiCase *ApiCase, result *ApiCaseResult) (flag bool) {
	flag = true
	asserts := apiCase.Asserts
	for _, assert := range asserts {
		fmt.Println(assert)
		assertResult := Assertion(&assert, result)
		result.ResponseAsserts = append(result.ResponseAsserts, assertResult)
		if !assertResult.Result {
			flag = false
		}
	}
	return
}

func Assertion(assertion *Asserts, result *ApiCaseResult) (assertResult AssertsResult) {
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
	return
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

// 比较相对

// 字符串正则替换变量
func replaceKeyFromMap(str string, vars map[string]interface{}) (result string) {
	for key, _ := range vars {
		re3, _ := regexp.Compile("\\$\\{" + key + "\\}")
		var temp interface{}
		if val, isOk := vars[key]; isOk {
			temp = val
		} else {
			fmt.Println("获取到key失败")
		}
		switch temp.(type) {
		case string:
			result = re3.ReplaceAllString(str, temp.(string))
		case int:
			result = re3.ReplaceAllString(str, fmt.Sprintf("%d", temp))
		case bool:
			result = re3.ReplaceAllString(str, fmt.Sprintf("%t", temp))
		case float64:
			result = re3.ReplaceAllString(str, fmt.Sprintf("%.2f", temp))
		}

	}
	fmt.Println(str)
	return
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
