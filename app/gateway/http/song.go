package http

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"strconv"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/app/song/repository/db/model"
	"web3-music-platform/idl/pb"
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

	response, err := rpc.Upload(ctx, &pb.CreateSongRequest{
		Song:    utils.ToSongModel(&song),
		Content: bytes,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
