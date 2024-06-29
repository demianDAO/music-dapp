package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"web3-music-platform/internal/app/gateway/services"
	"web3-music-platform/pkg/rdb"
	"web3-music-platform/pkg/utils"
)

func DiscoverySongs(ctx *gin.Context) {
	var logInstance = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "DiscoverySongs",
	})

	limitStr := ctx.Param("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 64)
	songs, err := services.DiscoverySongs(ctx, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	logInstance.Infof("songs:%v", songs)
	ctx.JSON(http.StatusOK, songs)
}

func UploadSong(ctx *gin.Context) {

	var logInstance = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "UploadSong",
	})

	fileHeader, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	title := ctx.PostForm("title")
	artistAddr := ctx.GetString("user_addr")
	overview := ctx.PostForm("overview")

	//tokenUri := ctx.PostForm("token_uri")
	tokenID, _ := strconv.ParseUint(ctx.PostForm("token_id"), 10, 64)
	//amount, _ := strconv.ParseUint(ctx.PostForm("amount"), 10, 64)
	//price, _ := strconv.ParseUint(ctx.PostForm("price"), 10, 64)

	logInstance.Infof("title:%v, artistAddr:%v, overview:%v tokenID:%v", title, artistAddr, overview, tokenID)

	go func() {
		fileBytes, err := utils.FileHeaderToBytes(fileHeader)
		err = services.UploadSong(ctx, title, artistAddr, overview, fileBytes, tokenID)
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
	var logInstance = log.WithFields(log.Fields{
		"module": "handlers",
		"func":   "FindSongs",
	})
	userAddr := ctx.Query("user_addr")

	songs, err := rdb.RedisInstance.GetSongs(ctx, userAddr)
	if err == redis.Nil || len(songs) == 0 {

		songs, err = services.FindSongs(ctx, userAddr)
		logInstance.Infof("songs:%v", songs)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		err = rdb.RedisInstance.SetSongs(ctx, userAddr, songs)

		if err != nil {
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

// turn to the front end
//func PurchaseSong(ctx *gin.Context) {
//
//	tokenId, _ := strconv.ParseUint(ctx.PostForm("token_id"), 10, 64)
//	singerAddr := ctx.PostForm("singer_addr")
//	userAddr := ctx.GetString("user_addr")
//
//	err := services.PurchaseSong(ctx, userAddr, singerAddr, tokenId)
//
//	if err != nil {
//		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"message": "purchase success",
//	})
//
//}
