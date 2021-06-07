package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Interface struct {
	Model

	Name       string `json:"name"`
	Method     string `json:"method"`
	Url        string `json:"url"`
	CreatedBy  int    `json:"created_by"`
	ModifiedBy int    `json:"modified_by"`
	ProjectId  int    `json:"project_id"`
	ModuleId   int    `json:"module_id"`
}

// InterfaceToUser 接口人员表, role 0 代表测试，1代表开发
type InterfaceToUser struct {
	Id          int `json:"id"`
	InterfaceId int `json:"interface_id"`
	UserId      int `json:"user_id"`
	Role        int `json:"role"`
}

type InterfaceAdd struct {
	Interface

	Develop []int `json:"develop"`
	Tester  []int `json:"tester"`
}

type InterfaceList struct {
	Interface
	UserName    string `json:"user_name"`
	ProjectName string `json:"project_name"`
	ModuleName  string `json:"module_name"`
}

type InterfaceUsersDetail struct {
	UserName string `json:"user_name"`
	UserId   string `json:"user_id"`
}

type InterfaceDetail struct {
	Interface

	ProjectName string                 `json:"project_name"`
	ModuleName  string                 `json:"module_name"`
	Develop     []InterfaceUsersDetail `json:"develop"`
	Tester      []InterfaceUsersDetail `json:"tester"`
}

// AddApi 添加接口
func AddApi(data *InterfaceAdd) error {
	tx := db.Begin()
	inface := &Interface{
		Name:      data.Name,
		Method:    data.Method,
		Url:       data.Url,
		CreatedBy: data.CreatedBy,
		ProjectId: data.ProjectId,
		ModuleId:  data.ModuleId,
	}
	if err := tx.Table("interface").Create(inface).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 存储开发人员
	for i := range data.Develop {
		err := tx.Table("interface_user").Create(&InterfaceToUser{
			InterfaceId: inface.Id,
			UserId:      data.Develop[i],
			Role:        1,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 存储测试人员
	for i := range data.Tester {
		err := tx.Table("interface_user").Create(&InterfaceToUser{
			InterfaceId: inface.Id,
			UserId:      data.Tester[i],
			Role:        0,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// CheckInterfaceNameExists 检查接口名称是否存在
func CheckInterfaceNameExists(name string, projectId int) (flag bool, err error) {
	var inter Interface
	err = db.Table("interface").Select("id").Where("name = ? and project_id = ?", name, projectId).First(&inter).Error
	if err != nil {
		return false, err
	}
	if inter.Id < 1 {
		return false, nil
	}
	return true, nil
}

// FindListByModuleId 根据模块id查询模块下的接口列表
func FindListByModuleId(pageSize, pageNum, moduleId int) (apiList []InterfaceList, count int, err error) {
	err = db.Offset(pageNum-1).Limit(pageSize).Raw("SELECT i.*,u.user_name,p.`name` AS project_name,m.NAME AS module_name FROM interface i LEFT JOIN `user` u"+
		" ON u.id =i.created_by LEFT JOIN project p ON p.id = i.project_id LEFT JOIN module m ON m.id = i.module_id WHERE "+
		"i.state = 1 and i.module_id = ? GROUP BY i.id DESC", moduleId).Scan(&apiList).Error
	db.Table("interface").Where("state = 1 and module_id = ?", moduleId).Count(&count)
	//err = db.Table("interface").Where("module_id = ? and state = 1", moduleId).Count(&count).Scan(&apiList).Error
	if err != nil {
		return nil, 0, err
	}
	return apiList, count, nil
}

// InterList 查询项目下的接口列表
func InterList(pageSize, pageNum, projectId int) (list []InterfaceList, count int, err error) {
	err = db.Offset(pageNum-1).Limit(pageSize).Raw("SELECT i.*,u.user_name,p.`name` AS project_name,m.NAME AS module_name FROM interface i LEFT JOIN `user` u"+
		" ON u.id =i.created_by LEFT JOIN project p ON p.id = i.project_id LEFT JOIN module m ON m.id = i.module_id WHERE "+
		"i.state = 1 and i.project_id = ? GROUP BY i.id DESC", projectId).Scan(&list).Error
	db.Table("interface").Where("state = 1 and project_id = ?", projectId).Count(&count)
	//err = db.Table("interface").Where("module_id = ? and state = 1", projectId).Count(&count).Scan(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

// 接口详情
func InterDetail(apiId int) (detail InterfaceDetail, err error) {
	err = db.Raw("select i.*, p.name as project_name, m.name as module_name from interface i left join project p on i.project_id = p.id left join module m"+
		" on i.module_id = m.id where i.id = ? and i.state = 1", apiId).Scan(&detail).Error

	err = db.Raw("select i.user_id, u.user_name from interface_user i left join user u on i.user_id = u.id where"+
		" i.role = 1 and interface_id = ?", apiId).Scan(&detail.Develop).Error

	err = db.Raw("select i.user_id, u.user_name from interface_user i left join user u on i.user_id = u.id where"+
		" i.role = 0 and interface_id = ?", apiId).Scan(&detail.Tester).Error
	return
}

// 修改接口
func InterUpdate(data *InterfaceAdd) error {
	tx := db.Begin()
	inter := &Interface{
		Name:      data.Name,
		Method:    data.Method,
		Url:       data.Url,
		CreatedBy: data.CreatedBy,
		ProjectId: data.ProjectId,
		ModuleId:  data.ModuleId,
	}
	if err := tx.Table("interface").Where("id = ?", data.Id).Update(inter).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除人员重新插入
	if err := tx.Table("interface_user").Where("project_id", data.ProjectId).Unscoped().Delete(inter).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 存储开发人员
	for i := range data.Develop {
		err := tx.Table("interface_user").Create(&InterfaceToUser{
			InterfaceId: inter.Id,
			UserId:      data.Develop[i],
			Role:        1,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// 存储测试人员
	for i := range data.Tester {
		err := tx.Table("interface_user").Create(&InterfaceToUser{
			InterfaceId: inter.Id,
			UserId:      data.Tester[i],
			Role:        0,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (project *Interface) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (project *Interface) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
