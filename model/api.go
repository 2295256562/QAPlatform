package model

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

// 添加接口
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

// 检查接口名称是否存在
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
