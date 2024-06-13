package main

import (
	"web3-music-platform/config"
	"web3-music-platform/internal/app/user/db"
	"web3-music-platform/internal/app/user/services"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/rdb"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	db.Init()
	rdb.Init()

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.UserName), // 微服务名字
		micro.Address(config.UserAddress),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(), services.NewUserService())
	// 启动微服务
	_ = microService.Run()
}
