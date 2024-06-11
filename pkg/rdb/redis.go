package rdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
	song_service_model "web3-music-platform/app/song/repository/db/model"
	user_service_model "web3-music-platform/app/user/repository/db/model"

	"web3-music-platform/config"
)

var RedisInstance *RedisClient

type RedisClient struct {
	*redis.Client
}

func Init() {

	log.Print("init redis")
	log.Print("config.RedisNetworkAddress", config.RedisNetworkAddress)
	log.Print("config.RedisPassword", config.RedisPassword)

	networkAddress := config.RedisNetworkAddress
	password := config.RedisPassword
	client := redis.NewClient(&redis.Options{
		Addr:     networkAddress,
		Password: password,
		DB:       0,
	})
	result, err := client.Ping(context.Background()).Result()
	log.Print("redis ping result:", result)
	if err != nil {
		panic(err)
	}
	RedisInstance = &RedisClient{}
	RedisInstance.Client = client
	//return &RedisClient{
	//	Client: client,
	//}
}

func (client *RedisClient) GetUserInfo(ctx context.Context, address string) (*user_service_model.User, error) {
	var user user_service_model.User
	key := "user:" + address + ":profile"
	bytes, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (client *RedisClient) SetUserInfo(ctx context.Context, address string, user *user_service_model.User) error {
	key := "user:" + address + ":profile"
	bytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := client.Set(ctx, key, bytes, 15*time.Minute).Err(); err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) StoreToken(ctx context.Context, address, token string) error {
	key := "user:" + address + ":token"
	return client.Set(ctx, key, token, 15*time.Minute).Err()
}

func (client *RedisClient) CheckToken(ctx context.Context, address string) (string, error) {
	key := "user:" + address + ":token"
	return client.Get(ctx, key).Result()
}

func (client RedisClient) AddSong(ctx context.Context, address string, song *song_service_model.Song) error {
	key := "song:" + address + ":songs"
	jsonData, err := json.Marshal(song)
	if err != nil {
		return err
	}
	err = client.LPush(ctx, key, jsonData).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client RedisClient) SetSongs(ctx context.Context, address string, songs []*song_service_model.Song) error {
	if len(songs) == 0 {
		return errors.New("song list is empty")
	}
	key := "song:" + address + ":songs"
	pipe := client.TxPipeline()
	for _, song := range songs {
		jsonData, err := json.Marshal(song)
		if err != nil {
			return err
		}
		pipe.LPush(ctx, key, jsonData)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetSongs 获取 Redis 中存储的歌曲列表
func (client RedisClient) GetSongs(ctx context.Context, address string) ([]*song_service_model.Song, error) {
	key := "song:" + address + ":songs"
	songStrings, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	songs := make([]*song_service_model.Song, 0, len(songStrings))
	for _, songString := range songStrings {
		var song *song_service_model.Song
		err := json.Unmarshal([]byte(songString), &song)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func (client *RedisClient) AddPost(ctx context.Context, address string, post *user_service_model.Post) error {
	key := "user:" + address + ":posts"
	jsonData, err := json.Marshal(post)
	if err != nil {
		return err
	}
	err = client.LPush(ctx, key, jsonData).Err()
	if err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) GetPosts(ctx context.Context, address string) ([]*user_service_model.Post, error) {
	key := "user:" + address + ":posts"

	jsonDataList, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var posts []*user_service_model.Post
	for _, jsonData := range jsonDataList {
		post := &user_service_model.Post{} // 为每个帖子分配内存空间
		err := json.Unmarshal([]byte(jsonData), post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
