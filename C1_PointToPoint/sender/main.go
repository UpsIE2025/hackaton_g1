package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	godotenv.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	r := gin.Default()

	r.POST("/", func(ctx *gin.Context) {
		var req struct {
			ID   int    `json:"id"`
			From string `json:"from"`
			To   string `json:"to"`
			Size string `json:"size"`
		}
		if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}
		data, _ := json.Marshal(&req)
		err := rdb.RPush(context.Background(), "parcels", string(data)).Err()
		if err != nil {
			slog.Error(err.Error())
		}
		ctx.Status(http.StatusOK)
	})

	r.Run(":80")
}
