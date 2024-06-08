package http

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"strconv"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/app/song/repository/db/model"
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

	err = rpc.Upload(ctx, song, bytes)

	if err != nil {
		log.Print("Upload Fail==>", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload success",
	})
}

func FindSongsHandler(ctx *gin.Context) {
	userAddr := ctx.Query("user_address")

	songs, err := rpc.FindSongs(ctx, userAddr)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
