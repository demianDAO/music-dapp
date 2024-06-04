package main

import (
	"fmt"
	"web3-music-platform/app/user/repository/cache"
	"web3-music-platform/app/user/repository/db/dao"
	"web3-music-platform/app/user/service"
	"web3-music-platform/config"
	"web3-music-platform/idl/pb"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.Init()
	cache.Init()

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterUserServiceHandler(microService.Server(), service.NewUserService())
	// 启动微服务
	_ = microService.Run()
}
