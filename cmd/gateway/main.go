package main

import (
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"web3-music-platform/app/gateway/router"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/config"
	"web3-music-platform/pkg/rdb"
)

func main() {
	config.Init()
	rpc.InitRPC()
	rdb.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)

	// 创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name(config.GatewayName),
		web.Address(config.GatewayAddress),
		// 将服务调用实例使用gin处理
		web.Handler(router.NewRouter()),
		web.Registry(etcdReg),
	)
	// 接收命令行参数
	_ = server.Init()
	_ = server.Run()
}
