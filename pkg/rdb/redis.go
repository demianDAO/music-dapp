package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
	"web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/config"
)

var RedisInstance *RedisClient

type RedisClient struct {
	*redis.Client
}

func Init() {
	networkAddress := config.RedisNetworkAddress
	//password := config.RedisPassword
	client := redis.NewClient(&redis.Options{
		Addr: networkAddress,
		//Password: password,
		DB: 0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	RedisInstance = &RedisClient{}
	RedisInstance.Client = client
	//return &RedisClient{
	//	Client: client,
	//}
}

func (client *RedisClient) GetUserInfo(ctx context.Context, address string) (model.User, error) {
	var user model.User
	bytes, err := client.Get(ctx, address).Bytes()
	if err != nil {
		return model.User{}, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (client *RedisClient) SetUserInfo(ctx context.Context, address string, user model.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := client.Set(ctx, address, bytes, time.Second*30).Err(); err != nil {
		return err
	}
	return nil
}
