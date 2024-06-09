package http

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/app/song/repository/db/model"
	"web3-music-platform/pkg/rdb"
	"web3-music-platform/pkg/utils"
)

func UploadSongHandler(ctx *gin.Context) {

	var song model.Song

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	bytes, err := utils.FileHeaderToBytes(fileHeader)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	song.Title = ctx.PostForm("title")
	song.ArtistAddress = ctx.PostForm("artist_address")
	song.Overview = ctx.PostForm("overview")
	song.NFTAddress = ctx.PostForm("nft_address")
	song.TokenID, _ = strconv.ParseUint(ctx.PostForm("token_id"), 10, 64)

	marshalIndent, err := json.MarshalIndent(&song, "", "  ")
	log.Print("UploadSongHandler==>", string(marshalIndent))

	// 将上传的歌曲信息添加到 Redis 缓存
	go func() {
		err = rdb.RedisInstance.AddSong(ctx, song.ArtistAddress, &song)
		if err != nil {
			log.Print("Failed to add song to Redis==>", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to cache song")
			return
		}
		err = rpc.Upload(ctx, song, bytes)
		if err != nil {
			log.Print("Upload Fail==>", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload success",
	})
}

func FindSongsHandler(ctx *gin.Context) {
	userAddr := ctx.Query("user_address")

	// 从 Redis 缓存中获取歌曲列表
	songs, err := rdb.RedisInstance.GetSongs(ctx, userAddr)
	log.Print("从 Redis 缓存中获取歌曲列表==>", songs)
	if err == redis.Nil || len(songs) == 0 {
		// 如果缓存中没有，则从数据库获取并缓存
		songs, err = rpc.FindSongs(ctx, userAddr)
		log.Print("缓存中没有，则从数据库获取并缓存==>", songs)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		err = rdb.RedisInstance.SetSongs(ctx, userAddr, songs)
		if err != nil {
			log.Print("Failed to add songs to Redis==>", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to cache songs")
			return
		}
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, songs)
}

func DownloadSongHandler(ctx *gin.Context) {
	txId := ctx.Param("tx_id")

	songs, err := rpc.DownloadSong(ctx, txId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	// 设置响应头
	ctx.Header("Content-Disposition", "attachment; filename=song.mp3")
	ctx.Header("Content-Type", "audio/mpeg")
	ctx.Header("Content-Length", string(len(songs)))

	// 将歌曲数据写入响应体
	ctx.Data(http.StatusOK, "audio/mpeg", songs)

	//ctx.JSON(http.StatusOK, songs)
}
