package main

import (
	"context"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"web3-music-platform/internal/app/song/db"
	"web3-music-platform/internal/app/song/repositories"
	"web3-music-platform/internal/app/song/services"
	"web3-music-platform/internal/irys"
	"web3-music-platform/internal/mq"
	"web3-music-platform/pkg/contract"

	"web3-music-platform/config"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/rdb"
)

func main() {
	config.Init()
	db.Init()
	rdb.Init()
	err := contract.NewGethClient()
	if err != nil {
		//log.Print("NewGethClient error ", err.Error())
		panic(err)
	}

	err = mq.Init(config.RabbitMqUrl)

	if err != nil {
		panic(err)
	}

	err = irys.NewIrysClient(config.PrivateKey, config.SepoliaRPC)
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
	_ = pb.RegisterSongServiceHandler(microService.Server(),
		services.NewSongService(contract.SongNFTTrade, repositories.NewSongRepository(context.Background()), contract.SongNFTTradeFilter, rdb.RedisInstance))
	// 启动微服务
	_ = microService.Run()
}
