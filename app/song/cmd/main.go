package main

import (
	"web3-music-platform/app/song/repository/cache"
	"web3-music-platform/app/song/repository/db/dao"
	"web3-music-platform/app/song/repository/db/decentralised_storage"
	"web3-music-platform/app/song/service"
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
	err := decentralised_storage.NewIrysClient(config.PrivateKey)
	if err != nil {
		panic(err)
	}

	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(config.SongName), // 微服务名字
		micro.Address(config.SongAddress),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = pb.RegisterSongServiceHandler(microService.Server(), service.NewSongService())
	// 启动微服务
	_ = microService.Run()
}
