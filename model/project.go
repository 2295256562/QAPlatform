package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Project struct {
	Model

	Name       string `json:"name"`
	Remake     string `json:"remake"`
	CreatedBy  int    `json:"created_by"`
	ModifiedBy int    `json:"modified_by"`
}

type AddProject struct {
	Project

	Member []int `json:"member"`
}

// 校验项目名称是否存在
func CheckProjectExist(name string) bool {
	var project Project
	err = db.Select("id").Where("name = ?", name).First(&project).Error
	if err != nil {
		return false
	}
	if project.Id < 1 {
		return false
	}
	return true
}

func CreateProject(data *AddProject) bool {
	err = db.Create(&Project{
		Name: data.Name,
	}).Error
	if err != nil {
		return false
	}

	db.Find()
	return true
}

func (project *Project) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (project *Project) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
