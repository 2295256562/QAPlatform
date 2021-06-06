package model

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"time"
)

type User struct {
	Model

	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

type UserIDAndUserName struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
}

// CheckUsernameExist 检查用户名是否存在
func CheckUsernameExist(username string) bool {
	var user User
	db.Select("id").Where("user_name = ?", username).First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

// CreateUser 创建用户
func CreateUser(data *User) bool {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return false
	}
	return true
}

// ListUser 查询用户列表
func ListUser(pageSize, pageNum int, maps interface{}) (users []User) {
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func Login(username, password string) (user User, err error) {
	err = db.Where("user_name = ?", username).First(&user).Error

	if err != nil {
		return user, errors.New(fmt.Sprint("查询不到此用户"))
	}
	if user.Password == ScryptPw(password) {
		return user, nil
	}
	return user, errors.New(fmt.Sprint("用户名或密码错误"))
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpm := base64.StdEncoding.EncodeToString(HashPw)
	return fpm
}

// UserListByRole 根据角色返回用户列表
func UserListByRole(role int) (users []UserIDAndUserName, err error) {
	if err := db.Table("user").Select("id, user_name").Where("role = ?", role).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedTime", time.Now().Unix())
	return nil
}

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedTime", time.Now().Unix())
	return nil
}
