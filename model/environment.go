package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Environment struct {
	Model

	Name       string `json:"name"`
	Domain     string `json:"domain"`
	Headers    string `json:"headers"`
	Variables  string `json:"variables"`
	ProjectId  int    `json:"project_id"`
	CreatedBy  int    `json:"created_by"`
	ModifiedBy int    `json:"modified_by"`
}

type Envs struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type EnvironmentList struct {
	Environment

	CreatedUser string `json:"created_user"`
}

// CheckNameExist 校验同项目中环境名称是否重复
func CheckNameExist(name string, projectId int) bool {
	var env Environment
	err := db.Table("environment").Select("id").Where("name = ? and project_id = ?").First(&env).Error

	if err != nil {
		return false
	}
	if env.Id < 0 {
		return false
	}
	return true
}

// AddEnvironment 添加环境
func AddEnvironment(data *Environment) (flag bool, err error) {
	err = db.Table("environment").Create(&Environment{
		Name:      data.Name,
		Domain:    data.Domain,
		Headers:   data.Headers,
		Variables: data.Variables,
		ProjectId: data.ProjectId,
		CreatedBy: data.CreatedBy,
	}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}

// Environments 项目列表(不带分页)
func Environments(projectId int) (envs []Envs, err error) {
	err = db.Table("environment").Select("id, name").Where("project_id = ? and state = 1", projectId).Scan(&envs).Error
	if err != nil {
		return nil, err
	}
	return envs, nil
}

// EnvironmentLists 项目带分页
func EnvironmentLists(pageSize, pageNum, projectId int) (envList []EnvironmentList, count int, err error) {

	err = db.Table("environment").Select("environment.*, user.user_name as created_user").
		Where("environment.state = 1 and environment.project_id = ?", projectId).Count(&count).
		Joins("left join user on user.id = environment.created_by").Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(&envList).Error
	if err != nil {
		return nil, 0, err
	}
	return
}

// EnvironmentDel 环境删除
func EnvironmentDel(envId int) bool {
	err := db.Table("environment").Where("id = ?", envId).Update("state", 0).Error

	if err != nil {
		return false
	}
	return true
}

// EnvironmentEdit 环境编辑
func EnvironmentEdit(data *Environment) (flag bool, err error) {
	err = db.Table("environment").Where("id = ?", data.Id).Updates(&Environment{
		Name:       data.Name,
		Domain:     data.Domain,
		Headers:    data.Headers,
		Variables:  data.Variables,
		ModifiedBy: data.ModifiedBy,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// EnvironmentDetail 项目详情
func EnvironmentDetail(id int) (env *EnvironmentList, err error) {
	err = db.Table("environment").Select("environment.*, user.user_name as created_user").
		Where("id = ?", id).Joins("left join user on user.id = environment.created_by").
		Scan(&env).Error
	if err != nil {
		return nil, err
	}
	return env, nil
}

func (env *Environment) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (env *Environment) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
