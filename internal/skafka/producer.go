package skafka

import "github.com/IBM/sarama"

// Producer ...
type Producer struct {
	p sarama.SyncProducer
}

// NewProducer ...
func NewProducer(kafkaAddr string) (*Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{kafkaAddr}, config)
	if err != nil {
		return nil, err
	}
	return &Producer{p: producer}, nil
}

// Close ...
func (p *Producer) Close() {
	p.Close()
}

// Send ...
func (p *Producer) Send(topic string, data string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}
	_, _, err := p.p.SendMessage(msg)
	return err
}
