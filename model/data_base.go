package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type DataBase struct {
	Model

	Name         string `json:"name"`
	DatabaseType string `json:"database_type"`
	Address      string `json:"address"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Remark       string `json:"remark"`
	EnvId        int    `json:"env_id"`
	ProjectId    int    `json:"project_id"`
	CreatedBy    int    `json:"created_by"`
	ModifiedBy   int    `json:"modified_by"`
}

type DataBaseList struct {
	DataBase

	EnvName     string `json:"env_name"`
	CreatedUser string `json:"created_user"`
}

func AddDataBase(data *DataBase) (err error) {
	if err = db.Debug().Table("dataBase").Create(&data).Error; err != nil {
		return err
	}
	return
}

func EditDataBase(data *DataBase) (err error) {
	err = db.Debug().Table("dataBase").Where("id = ?", data.Id).Updates(data).Error
	if err != nil {
		return
	}
	return
}

func DetailDataBase(id int) (detail DataBase, err error) {
	err = db.Debug().Table("dataBase").Where("id = ?", id).Scan(&detail).Error
	if err != nil {
		return
	}
	return
}

func DataBaseLists(pageSize, pageNum, projectId int) (DataBaseList []DataBaseList, count int, err error) {
	err = db.Table("`dataBase` as d").Select("d.*, u.user_name as created_user, e.name as env_name").
		Joins("left join user u on d.created_by = u.id left join environment e on d.env_id = e.id where d.project_id = ? and d.state = 1", projectId).
		Count(&count).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&DataBaseList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func DataBaseDelete(id int) (err error) {
	err = db.Debug().Table("dataBase").Where("id = ?", id).Update("state", 0).Error
	if err != nil {
		return err
	}
	return
}

func (db *DataBase) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (db *DataBase) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
