package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	defer rdb.Close()

	r := gin.Default()

	r.POST("/publish", func(ctx *gin.Context) {
		var req struct {
			Message string `json:"message"`
		}
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		ctx.Status(http.StatusOK)
	})

	r.Run(":80")
}
