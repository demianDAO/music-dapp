package rdb

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

func (client *RedisClient) GetUserInfo(ctx context.Context, address string) (*model.User, error) {
	var user model.User
	bytes, err := client.Get(ctx, address).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (client *RedisClient) SetUserInfo(ctx context.Context, address string, user *model.User) error {
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := client.Set(ctx, address, bytes, 15*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) StoreToken(ctx context.Context, address, token string) error {
	return client.Set(ctx, address+":token", token, 15*time.Minute).Err()
}

func (client *RedisClient) CheckToken(ctx context.Context, address string) (string, error) {
	return client.Get(ctx, address+":token").Result()
}

func (client *RedisClient) AddPost(ctx context.Context, address string, post *model.Post) error {
	jsonData, err := json.Marshal(post)
	if err != nil {
		return err
	}
	err = client.LPush(ctx, address+":posts", jsonData).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) GetPosts(ctx context.Context, address string) ([]*model.Post, error) {
	jsonDataList, err := client.LRange(ctx, address+":posts", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var posts []*model.Post
	for _, jsonData := range jsonDataList {
		post := &model.Post{} // 为每个帖子分配内存空间
		err := json.Unmarshal([]byte(jsonData), post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
