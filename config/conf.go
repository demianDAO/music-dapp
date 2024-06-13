package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	DSN string

	EtcdAddress string

	RedisNetworkAddress string
	RedisPassword       string
	RedisDbName         int

	PrivateKey string
	SepoliaRPC string

	//	server
	GatewayName    string
	GatewayAddress string

	UserName    string
	UserAddress string

	SongName    string
	SongAddress string

	RabbitMqUrl string

	NftAddr    string
	NftMgrAddr string

	TestAddr1    string
	TestAddr2    string
	TestPrivate1 string
	TestPrivate2 string
)

func Init() {

	env := os.Getenv("APP_ENV")
	var configFilePath string
	log.Print("env:", env)
	switch env {
	case "prod":
		configFilePath = "config/config_prod.ini"
	case "test":
		configFilePath = "config/config_test.ini"
	case "dev":
		configFilePath = "config/config_dev.ini"
	default:
		log.Fatalf("Unknown environment: %s", env)
	}

	file, err := ini.Load(configFilePath)

	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadMysqlData(file)
	LoadEtcd(file)
	LoadRedisData(file)
	LoadIrys(file)
	LoadService(file)
	LoadRabbitMq(file)
	LoadSmartContract(file)
}

func LoadMysqlData(file *ini.File) {
	DSN = file.Section("mysql").Key("DSN").String()
}

func LoadEtcd(file *ini.File) {
	EtcdAddress = file.Section("etcd").Key("EtcdAddress").String()
}

func LoadRedisData(file *ini.File) {
	RedisNetworkAddress = file.Section("redis").Key("RedisAddr").String()
	RedisPassword = file.Section("redis").Key("RedisPassword").String()
}

func LoadIrys(file *ini.File) {
	PrivateKey = file.Section("irys").Key("PrivateKey").String()
	SepoliaRPC = file.Section("irys").Key("SepoliaRPC").String()

}

func LoadService(file *ini.File) {
	GatewayName = file.Section("gateway_service").Key("Name").String()
	GatewayAddress = file.Section("gateway_service").Key("Address").String()

	UserName = file.Section("user_service").Key("Name").String()
	UserAddress = file.Section("user_service").Key("Address").String()

	SongName = file.Section("song_service").Key("Name").String()
	SongAddress = file.Section("song_service").Key("Address").String()

}

func LoadRabbitMq(file *ini.File) {
	RabbitMqUrl = file.Section("rabbitmq").Key("RabbitMqUrl").String()
}

func LoadSmartContract(file *ini.File) {

	NftAddr = file.Section("smart_contract").Key("NftAddr").String()
	NftMgrAddr = file.Section("smart_contract").Key("NftMgrAddr").String()

	TestAddr1 = file.Section("smart_contract").Key("TestAddr1").String()
	TestAddr2 = file.Section("smart_contract").Key("TestAddr2").String()

	TestPrivate1 = file.Section("smart_contract").Key("TestPrivate1").String()
	TestPrivate2 = file.Section("smart_contract").Key("TestPrivate2").String()

}
