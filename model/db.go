package model

import (
	"QAPlatform/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB
var err error

type Model struct {
	Id           int `gorm:"primary_key" json:"id"`
	CreatedTime  int `json:"created_time"`
	ModifiedTime int `json:"modified_time"`
}

func InitDb() {
	fmt.Println("用户名 = ", utils.DbUser)
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	if err != nil {
		fmt.Println("连接数据库失败，请检查配置", err)
	}
	db.SingularTable(true)
	db.AutoMigrate()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
}
