package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_type "web3-music-platform/internal/api/handlers/type"
	"web3-music-platform/internal/api/models"
	"web3-music-platform/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	addr := ctx.Query("addr")
	user, err := h.service.GetUserByAddr(ctx, addr)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var params _type.CreateUserRequest
	ctx.BindJSON(params)
	user := models.User{
		Address: params.Address,
		Name:    params.Name,
	}

	err := h.service.CreateUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, user)
}
