package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	dbzAddr := fmt.Sprintf("%s:%s", os.Getenv("DEBEZIUM_HOST"), os.Getenv("DEBEZIUM_PORT"))
	kafkaAddr := fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))

	configureDbzConnector(dbzAddr)

	db, err := replicaDBInit()
	if err != nil {
		panic(err)
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{kafkaAddr}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	const topic = "postdb.public.posts"
	partition, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partition.Close()

	slog.Info("CDC Consumer started!")

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partition.Errors():
				slog.Error(err.Error())
			case msg := <-partition.Messages():
				handleEvent(db, msg)
			case <-sigchan:
				slog.Info("Gracefully shutdown!")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
}
