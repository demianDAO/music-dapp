package http

import (
	"net/http"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		//ctx.JSON(http.StatusBadRequest, cache.RespError(ctx, err, "UserRegister Bind 绑定参数失败"))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.Bind(&req); err != nil {
		//ctx.JSON(http.StatusBadRequest, cache.RespError(ctx, err, "UserLogin Bind 绑定参数失败"))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, cache.RespError(ctx, err, "UserLogin RPC 调用失败"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}
	token, err := utils.GenerateToken(userResp.UserDetail.Id)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, cache.RespError(ctx, err, "GenerateToken 失败"))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  userResp,
	})
}
