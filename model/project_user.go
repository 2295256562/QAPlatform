package model

type ProjectToUser struct {
	UserId    int `json:"user_id"`
	ProjectId int `json:"project_id"`
}
