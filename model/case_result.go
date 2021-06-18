package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ApiCaseResultStr struct {
	Id                 int    `json:"id"`
	CaseName           string `json:"case_name"`
	CaseId             int    `json:"case_id"`
	InterfaceId        int    `json:"interface_id"`
	EnvName            string `json:"env_name"`
	SuiteId            int    `json:"suite_id"`
	Method             string `json:"method"`
	Url                string `json:"url"`
	ResultType         int    `json:"result_type"`
	RequestHeaders     string `json:"request_headers"`
	RequestQuery       string `json:"request_query"`
	RequestBodyType    string `json:"request_body_type"`
	RequestBody        string `json:"request_body"`
	ResponseStatusCode int    `json:"response_status_code"`
	ResponseBody       string `json:"response_body"`
	ResponseHeaders    string `json:"response_headers"`
	ResponseTime       int    `json:"response_time"`

	ResponseAsserts  string `json:"response_asserts"`
	ResponseExtracts string `json:"response_extracts"`
	Exception        string `json:"exception"`
	ProjectId        int    `json:"project_id"`
	CreatedBy        int    `json:"created_by"`
	CreatedTime      int    `json:"created_time"`
}

func AddCaseResult(result *ApiCaseResultStr) (id int, err error) {
	err = db.Debug().Table("case_result").Create(&result).Error
	if err != nil {
		return 0, err
	}
	return result.Id, nil
}

func QueryCaseResult(id int) (result ApiCaseResultStr, err error) {
	err = db.Debug().Table("case_result").Where("id = ?", id).Scan(&result).Error
	if err != nil {
		return
	}
	return result, nil
}

func (env *ApiCaseResultStr) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}
