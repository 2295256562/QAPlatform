package utils

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	//file, err := ini.Load("config/config.ini")
	envFile := "config/config.ini"

	for i := 0; i < 5; i++ {
		if _, err := os.Stat(envFile); err == nil {
			break
		} else {
			envFile = "../" + envFile
		}
	}

	file, err := ini.Load(envFile)
	if err != nil {
		fmt.Println("加载配置文件错误，请检查配置文件")
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("113.31.147.158")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("peishan")
	DbPassword = file.Section("database").Key("DbPassword").MustString("1234")
	DbName = file.Section("database").Key("DbName").MustString("QAPlatform")
}
