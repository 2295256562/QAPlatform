package model

import (
	"QAPlatform/utils"
	"github.com/jinzhu/gorm"
	"time"
)

type InterfaceCase struct {
	Model

	Name        string `json:"name"`
	Type        string `json:"type"`
	Parameters  string `json:"parameters"`
	Headers     string `json:"headers"`
	Query       string `json:"query"`
	Asserts     string `json:"asserts"`
	Extract     string `json:"extract"`
	Remark      string `json:"remark"`
	InterfaceId int    `json:"interface_id"`
	EnvId       int    `json:"env_id"`

	CreatedBy  int `json:"created_by"`
	ModifiedBy int `json:"modified_by"`
	ProjectId  int `json:"project_id"`
}

type InterfaceQueryDto struct {
	Name        string `form:"name" json:"name"`
	ProjectId   int    `form:"project_id" json:"project_id"`
	EnvId       int    `form:"env_id" json:"env_id"`
	InterfaceId int    `form:"interface_id" json:"interface_id"`
	CreatedBy   int    `form:"created_by" json:"created_by"`
	Page        int    `form:"page" json:"page"`
	PageSize    int    `form:"page_size" json:"page_size"`
}

type InterCaseList struct {
	InterfaceCase

	InterfaceName string `json:"interface_name"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	Domain        string `json:"domain"`
	CreatedUser   string `json:"created_user"`
}

// InterfaceCaseAdd 添加接口测试用例
func InterfaceCaseAdd(data *InterfaceCase) (err error) {
	if db.Table("interface_case").Create(&data).Error != nil {
		return err
	}
	return
}

// CaseList 查询接口列表
func CaseList(data *InterfaceQueryDto) (InterCaseList []InterCaseList) {
	tx := db.Debug().Table("interface_case")
	tx = tx.Select("interface_case.*, interface.name as interface_name, interface.url, interface.method, environment.domain, user.user_name as created_user")
	tx = tx.Where("interface_case.project_id = ? and interface_case.state = 1", data.ProjectId)

	if data.Name != "" {
		tx = tx.Where("name = ?", data.Name)
	}
	if data.EnvId > 0 {
		tx = tx.Where("env_id = ?", data.EnvId)
	}

	if data.InterfaceId > 0 {
		tx = tx.Where("interface_id = ?", data.InterfaceId)
	}

	if data.CreatedBy > 0 {
		tx.Where("created_by = ?", data.CreatedBy)
	}
	if data.Page > 0 && data.PageSize > 0 {
		tx = tx.Limit(data.PageSize).Offset((data.Page - 1) * data.PageSize)
	}

	tx = tx.Joins("left join interface on interface.id = interface_case.interface_id" +
		" left join environment on environment.id = interface_case.env_id left join user on interface_case.created_by = user.id")
	tx.Find(&InterCaseList).RecordNotFound()
	return InterCaseList
}

// CaseDetail 用例详情
func CaseDetail(id int) (caseDetail InterfaceCase, err error) {
	if err = db.Table("interface_case").Where("id = ?", id).Find(&caseDetail).Error; err != nil {
		return
	}
	return
}

// 用例执行
func CaseInfo(id int) (apiCase utils.ApiCaseStr, err error) {

	if err = db.Debug().Table("interface_case as c").
		Select("c.*, e.domain, e.variables as g_vars, e.headers as env_headers, e.name as env_name, i.name as interface_name, i.url, i.method").
		Joins("left join environment e on c.env_id = e.id left join interface i on c.interface_id = i.id").
		Where("c.id = ?", id).Scan(&apiCase).Error; err != nil {
		return
	}
	return
}

func CaseEdit(data *InterfaceCase) (err error) {
	err = db.Table("interface_case").Where("id = ?", data.Id).Updates(&data).Error
	if err != nil {
		return
	}
	return
}

func (interCase *InterfaceCase) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (interCase *InterfaceCase) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
