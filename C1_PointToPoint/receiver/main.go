package main

import (
	"context"
	"fmt"
	"time"

	"log/slog"

	"github.com/redis/go-redis/v9"
)

const channel = "parcels"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	ctx := context.Background()

	slog.Info("Waiting for parcels...")
	for {
		r := rdb.BRPop(ctx, 1*time.Second, channel)
		if len(r.Val()) > 0 {
			fmt.Println(r.Val())
		}
	}
}
