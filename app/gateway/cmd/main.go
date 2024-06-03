package main

import (
	"fmt"
	"web3-music-platform/app/gateway/router"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/app/user/repository/cache"
	"web3-music-platform/internal/config"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	config.Init()
	rpc.InitRPC()
	cache.InitCache()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		web.Registry(etcdReg),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
