package skafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

// Consumer ...
type Consumer struct {
	C sarama.Consumer
}

// NewConsumer ...
func NewConsumer(kafkaAddr string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// Create new consumer
	consumer, err := sarama.NewConsumer([]string{kafkaAddr}, config)
	if err != nil {
		return nil, err
	}
	return &Consumer{C: consumer}, nil
}

// Close ...
func (c *Consumer) Close() {
	c.C.Close()
}

// Subscribe ...
func (c *Consumer) Subscribe(topic string, handler func(m *sarama.ConsumerMessage)) error {
	partition, err := c.C.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		return err
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-partition.Errors():
				fmt.Println(err)
			case msg := <-partition.Messages():
				handler(msg)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	return err
}
