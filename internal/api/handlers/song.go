package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_type "web3-music-platform/internal/api/handlers/type"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/service"
)

type SongHandler struct {
	service *service.SongService
}

func NewSongHandler(_service *service.SongService) *SongHandler {
	return &SongHandler{
		service: _service,
	}
}

func (sh SongHandler) CreateSongHandler(ctx *gin.Context) {
	var params _type.SongCreateRequest
	ctx.BindJSON(params)
	song := models.Song{
		Title:      params.Title,
		AuthorAddr: params.AuthorAddr,
		Overview:   params.Overview,
	}

	err := sh.service.CreateSong(ctx, song)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, song)
}

func (sh SongHandler) GetAllSongByAuthorHandler(ctx *gin.Context) {
	authorAddr := ctx.Query("authorAddr")

	songs, err := sh.service.GetAllSongByAuthor(ctx, authorAddr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, songs)
}
