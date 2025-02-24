package main

import (
	"context"
	"encoding/json"
	"fmt"

	"log/slog"

	"github.com/redis/go-redis/v9"
)

const channel = "decreaseStock"

type product struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type decreaseStockEvent struct {
	Products []product `json:"products"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	ctx := context.Background()

	s := rdb.Subscribe(ctx, channel)
	ch := s.Channel()

	slog.Info("Waiting for decreaseStockEvent...")
	for m := range ch {
		var event decreaseStockEvent
		err := json.Unmarshal([]byte(m.Payload), &event)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		slog.Info("Processing decreaseStockEvent from an order...")
		for _, p := range event.Products {
			slog.Info(fmt.Sprintf("== Decreasing stock for product %d in %d", p.ID, p.Quantity))
		}
		slog.Info("Finished ok")
	}

}
