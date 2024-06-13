package handlers

import (
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"web3-music-platform/internal/app/gateway/services"
	"web3-music-platform/pkg/grpc/pb"
	"web3-music-platform/pkg/rdb"
	"web3-music-platform/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var req pb.UserRegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResp, err := services.UserRegister(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	userAddr := ctx.Query("user_addr")
	//var user *models.User
	var token string

	user, err := rdb.RedisInstance.GetUserInfo(ctx, userAddr)
	log.Print("get user info from redis", user)

	if err == redis.Nil {
		user, err = services.UserLogin(ctx, userAddr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Print("set user info to redis", user)
		err = rdb.RedisInstance.SetUserInfo(ctx, userAddr, user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err = rdb.RedisInstance.CheckToken(ctx, userAddr)
	log.Print("get token from redis", token)
	if err == redis.Nil {
		log.Print("token in redis not found")
		token, err = utils.GenerateToken(user.Address)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Print("set token to redis", token)
		err := rdb.RedisInstance.StoreToken(ctx, userAddr, token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
