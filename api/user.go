package api

import (
	"QAPlatform/model"
	"QAPlatform/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func Login(c *gin.Context) {
	var user model.User
	checkParamsErr := c.ShouldBind(&user)

	if checkParamsErr != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("输入参数有误")))
		return
	}

	res, err := model.Login(user.UserName, user.Password)

	if err != nil {
		fmt.Println(err)
		utils.ResponseError(c, 500, errors.New(fmt.Sprint(err)))
		return
	} else {
		// 生成token
		token, err := utils.GenerateToken(res.Id, res.UserName)
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
}

func Register(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("输入参数有误")))
		return
	}

	usernameExist := model.CheckUsernameExist(user.UserName)
	if usernameExist {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用户已存在")))
		return
	}

	flag := model.CreateUser(&user)
	if !flag {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("用户注册失败")))
		return
	} else {
		utils.ResponseSuccess(c, "注册成功")
		return
	}
}

func QueryUserListOnRole(c *gin.Context) {
	role := com.StrTo(c.Query("role")).MustInt()

	users, err := model.UserListByRole(role)
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询用户列表错误")))
		return
	}

	if users == nil {
		arr := make([]int, 0)
		utils.ResponseSuccess(c, arr)
		return
	}
	utils.ResponseSuccess(c, users)
	return
}

func UserAll(c *gin.Context) {
	users, err := model.Users()
	if err != nil {
		utils.ResponseError(c, 500, errors.New(fmt.Sprint("查询用户列表错误")))
		return
	}

	if users == nil {
		arr := make([]int, 0)
		utils.ResponseSuccess(c, arr)
		return
	}
	utils.ResponseSuccess(c, users)
	return
}
