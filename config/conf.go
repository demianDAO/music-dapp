package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	DSN string

	EtcdAddress string

	RedisNetworkAddress string
	RedisPassword       string
	RedisDbName         int

	PrivateKey string

	//	server
	GatewayName    string
	GatewayAddress string

	UserName    string
	UserAddress string

	SongName    string
	SongAddress string
)

func Init() {
	file, err := ini.Load("./config_dev.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadRedisData(file)
	LoadIrys(file)
	LoadService(file)
}

func LoadMysqlData(file *ini.File) {
	DSN = file.Section("mysql").Key("DSN").String()
}

func LoadEtcd(file *ini.File) {
	EtcdAddress = file.Section("etcd").Key("EtcdAddress").String()
}

func LoadRedisData(file *ini.File) {
	RedisNetworkAddress = file.Section("cache").Key("RedisNetworkAddress").String()
	RedisPassword = file.Section("cache").Key("RedisPassword").String()
}

func LoadIrys(file *ini.File) {
	PrivateKey = file.Section("irys").Key("PrivateKey").String()
}

func LoadService(file *ini.File) {
	GatewayName = file.Section("gateway_service").Key("Name").String()
	GatewayAddress = file.Section("gateway_service").Key("Address").String()

	UserName = file.Section("user_service").Key("Name").String()
	UserAddress = file.Section("user_service").Key("Address").String()

	SongName = file.Section("song_service").Key("Name").String()
	SongAddress = file.Section("song_service").Key("Address").String()

}
