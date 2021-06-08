package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Module struct {
	Model

	Name       string `json:"name"`
	ProjectId  int    `json:"project_id"`
	CreatedBy  int    `json:"created_by"`
	ModifiedBy int    `json:"modified_by"`
}

type ModuleDetail struct {
	Model

	Name        string `json:"name"`
	ProjectId   int    `json:"project_id"`
	CreatedBy   int    `json:"created_by"`
	CreatedName string `json:"created_name"`
	ModifiedBy  int    `json:"modified_by"`
}

func CheckModuleNameExist(name string, projectId int) bool {
	var module Module
	err = db.Select("id").Where("name = ? and project_id = ?", name, projectId).First(&module).Error
	if err != nil {
		return false
	}
	if module.Id < 1 {
		return false
	}
	return true
}

func AddModule(data *Module) bool {
	err := db.Table("module").Create(&Module{
		Name:      data.Name,
		ProjectId: data.ProjectId,
		CreatedBy: data.CreatedBy,
	}).Error

	if err != nil {
		return false
	}
	return true
}

func GetModuleList(pageSize, pageNum int, maps interface{}) (modules []ModuleDetail, count int) {
	err := db.Table("module").Select("module.*, user.user_name as created_name").Joins("left join user on user.id = " +
		"module.created_by").
		Where(maps).Count(&count).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&modules).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func GetModules(projectId int) (ModuleDetails []ModuleDetail, err error) {
	err = db.Debug().Table("module").Select("module.*, user.user_name as created_name").Joins("left join user on user.id = "+
		"module.created_by").Where("project_id = ? and module.state = 1", projectId).Scan(&ModuleDetails).Error
	if err != nil {
		return nil, err
	}
	return ModuleDetails, nil
}

func ModuleDel(id int) bool {
	err := db.Table("module").Where("id = ?", id).Update("state", 0).Error

	if err != nil {
		return false
	}
	return true
}

func ModuleEdit(data *Module) bool {
	data.ModifiedTime = int(time.Now().Unix())
	err := db.Table("module").Where("id = ?", data.Id).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

func (module *Module) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (module *Module) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())

	return nil
}
