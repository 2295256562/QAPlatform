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
	Members []int `json:"members"`
}

type ProjectUserList struct {
	UserName string `json:"user_name"`
	UserId   int    `json:"user_id"`
}

type ProjectDetail struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Remake  string            `json:"remake"`
	Members []ProjectUserList `json:"members"`
}

// CheckProjectExist 校验项目名称是否存在
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
	project := &Project{
		Name:      data.Name,
		Remake:    data.Remake,
		CreatedBy: data.CreatedBy,
	}
	err = db.Create(project).Error
	if err != nil {
		return false
	}
	for i := range data.Members {
		db.Table("project_user").Create(&ProjectToUser{
			UserId:    data.Members[i],
			ProjectId: project.Id,
		})
	}
	return true
}

func GetProjectList(pageSize, pageNum int, maps interface{}) (projects []Project, count int) {
	err := db.Table("project").Where(maps).Count(&count).Offset(pageNum - 1).Limit(pageSize).Find(&projects).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func GetProjectDetail(id int) (projectDetail ProjectDetail) {
	db.Table("project").Where("id = ?", id).First(&projectDetail)
	db.Raw("SELECT `user`.id as user_id, user.user_name from project_user LEFT JOIN user on project_user.user_id = `user`.id WHERE project_user.project_id = ? ORDER BY project_user.id", id).Scan(&projectDetail.Members)
	//db.Table("project_user").Select("user_name", "user_id").Joins("LEFT JOIN user on project_user.user_id = `user`.id").Where("project_id = ?", id).First(&projectDetail.Members)
	//db.Table("project").Model(&projectDetail).Select("id = ?", id)
	//db.Table("project_user").Select("project_id = ?", id).Scan(&projectDetail.Members)
	//db.Raw("SELECT p.id, p.`name`, p.remake, t.user_name, t.id as user_id FROM (SELECT pu.project_id, u.user_name, "+
	//	"u.id FROM project_user pu RIGHT JOIN `user` u on  u.id = pu.user_id)"+
	//	" t LEFT JOIN project p on t.project_id = p.id WHERE p.id = ?", id).Scan(&projectDetail)
	return
}

func (project *Project) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (project *Project) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
