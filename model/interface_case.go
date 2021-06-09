package model

import (
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

// 添加接口测试用例
func InterfaceCaseAdd(data *InterfaceCase) (err error) {
	if db.Table("interface_case").Create(&data).Error != nil {
		return err
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
