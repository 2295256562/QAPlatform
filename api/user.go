package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	json := model.User{}
	user, flag := model.Login(json.UserName, json.Password)
	if !flag {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("登录失败，用户名或密码错误")))
		return
	}
	// 生成token
	token, err := utils.GenerateToken(user.Id, user.UserName)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("生成token错误")))
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	data["user_name"] = user.UserName
	utils.ResponseSuccess(c, data)
	return
}
