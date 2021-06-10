package utils

import (
	"encoding/json"
	"fmt"
	"github.com/asmcos/requests"
	"github.com/oliveagle/jsonpath"
	"regexp"
	"time"
)

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

	ResponseAsserts  string `json:"response_asserts"`
	ResponseExtracts string `json:"response_extracts"`
	Exception        string `json:"exception"`
	ProjectId        int    `json:"project_id"`
	CreatedBy        int    `json:"created_by"`
	ModifiedBy       int    `json:"modified_by"`
	Id               int    `json:"id"`
	CreatedTime      int    `json:"created_time"`
	ModifiedTime     int    `json:"modified_time"`
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
func RequestExecutor(apiCase *ApiCase) (result ApiCaseResult) {
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

	//var err error
	//var resp requests.Response

	switch method {
	case "GET":
		resp, err := requests.Get(url + query)
		fmt.Println(resp.Text(), err)
	case "POST":
		parameters := apiCase.Parameters
		body := replaceKeyFromMap(parameters, apiCase.GVars)
		fmt.Println("body信息 ==》", body)
		if bodyType == "form" {
			mapData, _ := JsonToMap(body)
			resp, err := requests.Post(url, mapData)
			fmt.Println(resp.Text(), err)
		}
		if bodyType == "json" {
			var data map[string]interface{}
			if err := json.Unmarshal([]byte(body), &data); err == nil {
				resp, err := requests.PostJson(url, data)
				fmt.Println(resp.Text(), err)
			} else {
				fmt.Println("字符串转json失败：", err)
			}
		}
	}
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
//func Assertion(assertion Asserts) {
//	switch assertion.AssertType {
//	case "response_body":
//		JsonPathExtract()
//	}
//}

// jsonpath
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
