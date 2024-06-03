package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DSN string

	EtcdHost string
	EtcdPort string

	RedisNetworkAddress string
	RedisPassword       string
	RedisDbName         int
)

func Init() {
	file, err := ini.Load("./config_dev.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadRedisData(file)
}

func LoadMysqlData(file *ini.File) {
	DSN = file.Section("mysql").Key("DSN").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadRedisData(file *ini.File) {
	RedisNetworkAddress = file.Section("cache").Key("RedisNetworkAddress").String()
	RedisPassword = file.Section("cache").Key("RedisPassword").String()
}
