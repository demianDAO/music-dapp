package utils

import "github.com/gin-gonic/gin"

func handleError(ctx *gin.Context, statusCode int, err error) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}
