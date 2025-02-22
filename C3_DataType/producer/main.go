package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

// Función para crear un productor de Kafka
func createProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear productor Kafka: %v", err)
	}
	return producer, nil
}

// Función para enviar un mensaje a un topic
func sendMessage(producer sarama.SyncProducer, topic string, message string) {
	// Crear el mensaje
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Enviar el mensaje
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Error al enviar mensaje a %s: %v", topic, err)
		return
	}

	log.Printf("Mensaje enviado a partición %d con offset %d en el topic %s", partition, offset, topic)
}

func main() {
	// Crear un productor de Kafka
	producer, err := createProducer()
	if err != nil {
		log.Fatalf("Error al crear productor Kafka: %v", err)
	}
	defer producer.Close()

	// Datos de ejemplo de los sensores del vehículo
	vehicleData := map[string]string{
		"c3_fuel-topic":          "50",     // En porcentaje; Datos del sensor de gasolina
		"c3_speed-topic":         "80",     // En k,/h; Datos del sensor de velocidad
		"c3_engine-status-topic": "true",   // Datos del sensor de motor
		"c3_doors-status-topic":  "001101", // En binario para conocer puertas (5 puertas) Datos del sensor de puertas
	}

	// Enviar los datos de los sensores a sus respectivos topics
	for topic, message := range vehicleData {
		sendMessage(producer, topic, message)
	}

	// Agregar un pequeño delay para ver los resultados
	time.Sleep(2 * time.Second)
}
