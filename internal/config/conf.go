package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DSN string
)

func Init() {
	file, err := ini.Load("./config_dev.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
}

func LoadMysqlData(file *ini.File) {
	DSN = file.Section("mysql").Key("DSN").String()
}
