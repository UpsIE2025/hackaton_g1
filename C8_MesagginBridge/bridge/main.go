package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/redis/go-redis/v9"
)

type ecommerceData struct {
	OrderID  int `json:"id"`
	Products []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	} `json:"products"`
}

type inventarioDataProd struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type inventarioData struct {
	Products []inventarioDataProd `json:"products"`
}

const kafkaAddr = "localhost:9092"
const redisAddr = "localhost:6379"

const ecommerceChannel = "orders"
const inventarioChannel = "decreaseStock"

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	defer rdb.Close()

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{kafkaAddr}, config)
	if err != nil {
		panic(err)
	}

	par, err := consumer.ConsumePartition(ecommerceChannel, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)

	slog.Info("Waiting for order events")
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case msg := <-par.Messages():
				bridgeConvert(rdb, msg)
			case <-sigch:
				doneCh <- struct{}{}
				return
			}
		}
	}()
	<-doneCh
}

func bridgeConvert(rdb *redis.Client, msg *sarama.ConsumerMessage) {
	edata := ecommerceData{}
	err := json.Unmarshal(msg.Value, &edata)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info(fmt.Sprintf("Converting order %d", edata.OrderID))
	idata := inventarioData{}
	idata.Products = make([]inventarioDataProd, 0, len(edata.Products))
	for _, p := range edata.Products {
		idata.Products = append(idata.Products, inventarioDataProd{
			ID:       p.ID,
			Quantity: p.Quantity,
		})
	}
	dataToSend, _ := json.Marshal(idata)
	slog.Info("Sending descreaseStock event to inventario service")
	err = rdb.Publish(context.Background(), inventarioChannel, dataToSend).Err()
	if err != nil {
		slog.Error(err.Error())
	}
}
