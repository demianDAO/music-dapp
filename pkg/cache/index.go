package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
	"web3-music-platform/internal/config"
)

const (
	USER_ID_KEY string = "user_id_key"
)

type RedisClient struct {
	*redis.Client
}

func Init() *RedisClient {
	networkAddress := config.RedisNetworkAddress
	password := config.RedisPassword
	client := redis.NewClient(&redis.Options{
		Addr:     networkAddress,
		Password: password,
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	//RedisClient = client
	return &RedisClient{
		Client: client,
	}
}

func (client *RedisClient) GetUserInfo(ctx context.Context) (uint64, error) {
	curUserId, err := client.Get(ctx, USER_ID_KEY).Uint64()
	if err != nil {
		return 0, err
	}
	return curUserId, nil
}

func (client *RedisClient) SetUserInfo(ctx context.Context, userId uint64) error {
	if err := client.Set(ctx, USER_ID_KEY, userId, time.Second*30).Err(); err != nil {
		return err
	}
	return nil
}
