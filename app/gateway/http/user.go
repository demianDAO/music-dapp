package http

import (
	"log"
	"net/http"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	address := ctx.Query("address")
	var req pb.UserLoginRequest

	req.Address = address

	res, err := rpc.UserLogin(ctx, &req)
	log.Print("user login rpc response: ", res)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GenerateToken(res.User.Address)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  res.User,
	})
}
