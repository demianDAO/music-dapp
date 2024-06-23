package rdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
	user_service_model "web3-music-platform/internal/app/user/models"
	"web3-music-platform/pkg/grpc/pb"

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
}

func (client *RedisClient) CacheReleasedSongEvent(ctx context.Context, address string, tokenId uint64) error {
	address = strings.ToLower(address)
	key := "event:" + address + ":released song"
	if err := client.LPush(ctx, key, tokenId).Err(); err != nil {
		return err
	}
	return nil
}

func (client *RedisClient) GetOldestTokenId(ctx context.Context, address string) (uint64, error) {
	address = strings.ToLower(address)
	key := "event:" + address + ":released song"
	tokenIdStr, err := client.LPop(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.New("no tokenId found")
		}
		return 0, err
	}

	tokenId, err := strconv.ParseUint(tokenIdStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return tokenId, nil
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

func (client RedisClient) SetSongs(ctx context.Context, address string, songs []*pb.FindSongsByAddrRes_SongInfo) error {
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
	pipe.Expire(ctx, key, 15*time.Minute)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

// GetSongs 获取 Redis 中存储的歌曲列表
func (client RedisClient) GetSongs(ctx context.Context, address string) ([]*pb.FindSongsByAddrRes_SongInfo, error) {
	key := "song:" + address + ":songs"
	songStrings, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	songs := make([]*pb.FindSongsByAddrRes_SongInfo, 0, len(songStrings))
	for _, songString := range songStrings {
		var song *pb.FindSongsByAddrRes_SongInfo
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
