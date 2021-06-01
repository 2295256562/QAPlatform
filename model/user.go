package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	Model

	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// 检查用户名是否存在
func CheckUsernameExist(username string) bool {
	var user User
	db.Select("id").Where("name = ?", username).First(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

// 创建用户
func CreateUser(data *User) int {
	data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return 0
	}
	return 1
}

// 查询用户列表
func ListUser(pageSize, pageNum int, maps interface{}) (users []User) {
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	return
}

func Login(username, password string) {
	var user User
	db.Where("username = ?",username).First(&user)
	if user.Id == 0{
	}
}

// 密码加密
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
