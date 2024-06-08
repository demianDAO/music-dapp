package http

import (
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"web3-music-platform/app/gateway/rpc"
	"web3-music-platform/app/user/repository/db/model"
	"web3-music-platform/idl/pb"
	"web3-music-platform/pkg/rdb"
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
	var user *model.User
	var token string

	user, err := rdb.RedisInstance.GetUserInfo(ctx, address)
	log.Print("get user info from redis", user)

	if err == redis.Nil {
		res, err := rpc.UserLogin(ctx, &pb.UserLoginRequest{
			Address: address,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user = utils.ToDBUser(res.User)
		log.Print("set user info to redis", user)
		err = rdb.RedisInstance.SetUserInfo(ctx, address, user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err = rdb.RedisInstance.CheckToken(ctx, address)
	log.Print("get token from redis", token)
	if err == redis.Nil {
		log.Print("token in redis not found")
		token, err = utils.GenerateToken(user.Address)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Print("set token to redis", token)
		err := rdb.RedisInstance.StoreToken(ctx, address, token)
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
