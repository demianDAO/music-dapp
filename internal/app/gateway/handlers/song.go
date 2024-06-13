package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strconv"
	"web3-music-platform/internal/app/gateway/services"
	"web3-music-platform/internal/app/song/models"
	"web3-music-platform/pkg/rdb"
	"web3-music-platform/pkg/utils"
)

func UploadSong(ctx *gin.Context) {

	var song models.Song

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	song.Title = ctx.PostForm("title")
	song.ArtistAddr = ctx.PostForm("artist_addr")
	song.Overview = ctx.PostForm("overview")
	song.NFTAddr = ctx.PostForm("nft_addr")
	song.TokenID, _ = strconv.ParseUint(ctx.PostForm("token_id"), 10, 64)
	tokenUri := ctx.PostForm("token_uri")
	userAddr, exists := ctx.Get("user_addr")
	if !exists {
		panic("user_addr not found in context")
	}
	userAddrStr, ok := userAddr.(string)

	if !ok {
		panic("user_addr is not a string")
	}

	song.UserAddr = userAddrStr

	marshalIndent, err := json.MarshalIndent(&song, "", "  ")
	log.Print("UploadSong==>", string(marshalIndent))

	go func() {
		err = rdb.RedisInstance.AddSong(ctx, song.UserAddr, &song)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to cache song")
			return
		}
		fileBytes, err := utils.FileHeaderToBytes(fileHeader)
		err = services.UploadSong(ctx, song, fileBytes, tokenUri)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
	}()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload success",
	})
}

func FindSongs(ctx *gin.Context) {
	userAddr := ctx.Query("user_address")

	songs, err := rdb.RedisInstance.GetSongs(ctx, userAddr)
	//todo:tx_id更新后，没有同步到redis
	if err == redis.Nil || len(songs) == 0 {

		songs, err = services.FindSongs(ctx, userAddr)

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

func DownloadSong(ctx *gin.Context) {
	txId := ctx.Param("tx_id")

	songs, err := services.DownloadSong(ctx, txId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Header("Content-Disposition", "attachment; filename=song.mp3")
	ctx.Header("Content-Type", "audio/mpeg")
	ctx.Header("Content-Length", string(len(songs)))
	ctx.Data(http.StatusOK, "audio/mpeg", songs)
}
