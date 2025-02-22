package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

// Función para crear un consumidor de Kafka
func createConsumer() (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear consumidor Kafka: %v", err)
	}
	return consumer, nil
}

// Función para consumir mensajes de un topic
func consumeMessages(consumer sarama.Consumer, topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error al consumir mensajes del topic %s: %v", topic, err)
	}
	defer partitionConsumer.Close()

	// Leer mensajes
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Mensaje recibido en topic %s: %s\n", topic, string(msg.Value))
	}
}

func consumeMessagesv2(consumer sarama.Consumer, topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Error al consumir mensajes del topic %s: %v", topic, err)
	}
	defer partitionConsumer.Close()

	// Leer mensajes
	for msg := range partitionConsumer.Messages() {
		// Mostrar mensaje crudo recibido
		fmt.Printf("Mensaje recibido en topic %s: %s\n", topic, string(msg.Value))

		// Verificar si el mensaje es válido según el tipo esperado (int para speed-topic)
		if topic == "c3_speed-topic" {
			// Intentar convertir el mensaje a int
			speed, err := strconv.Atoi(string(msg.Value))
			if err != nil {
				// Si la conversión falla, mostrar un error
				log.Printf("Error: El valor recibido en %s no es una Velocidad Valida.", topic)
				continue // Continuar al siguiente mensaje si el valor no es válido
			}

			// Si la conversión es exitosa, mostrar el valor
			fmt.Printf("Velocidad recibida: %d km/h\n", speed)
		}
	}
}

func main() {
	// Crear un consumidor de Kafka
	consumer, err := createConsumer()
	if err != nil {
		log.Fatalf("Error al crear consumidor Kafka: %v", err)
	}
	defer consumer.Close()
	log.Print("Inicio de micro consumidor")

	// Consumir mensajes de los topics
	//go consumeMessages(consumer, "c3_fuel-topic")
	//go consumeMessages(consumer, "c3_speed-topic")
	//go consumeMessages(consumer, "c3_engine-status-topic")
	//go consumeMessages(consumer, "c3_doors-status-topic")
	go consumeMessagesv2(consumer, "c3_fuel-topic")
	go consumeMessagesv2(consumer, "c3_speed-topic")
	go consumeMessagesv2(consumer, "c3_engine-status-topic")
	go consumeMessagesv2(consumer, "c3_doors-status-topic")

	// Mantener el programa en ejecución
	select {}
}
