package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

// Estructura para recibir el JSON desde el cliente (Postman)
type VehicleData struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

// Crear productor de Kafka
func createProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		return nil, fmt.Errorf("no se pudo crear productor Kafka: %v", err)
	}
	return producer, nil
}

// Función para enviar el mensaje a Kafka
func sendMessageToKafka(topic string, message string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Enviar el mensaje
	_, _, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("error al enviar mensaje a Kafka: %v", err)
	}

	return nil
}

// Manejador para la solicitud POST
func sendData(w http.ResponseWriter, r *http.Request) {
	// Leer el cuerpo de la solicitud
	var vehicleData VehicleData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&vehicleData)
	if err != nil {
		http.Error(w, "Error al leer los datos de la solicitud", http.StatusBadRequest)
		return
	}

	// Validar si el topic es "speed-topic" y si el mensaje puede convertirse a int
	if vehicleData.Topic == "speed-topic" {
		_, err := strconv.Atoi(vehicleData.Message)
		if err != nil {
			http.Error(w, "El valor de velocidad debe ser un número entero", http.StatusBadRequest)
			return
		}
	}

	// Enviar mensaje a Kafka
	err = sendMessageToKafka(vehicleData.Topic, vehicleData.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al enviar mensaje a Kafka: %v", err), http.StatusInternalServerError)
		return
	}

	// Responder con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensaje enviado a Kafka con éxito"))
}

func main() {
	// Crear productor de Kafka
	var err error
	producer, err = createProducer()
	if err != nil {
		log.Fatalf("Error al crear productor Kafka: %v", err)
	}
	defer producer.Close()

	// Crear el servidor HTTP
	http.HandleFunc("/send", sendData) // Endpoint POST /send
	log.Println("Servidor escuchando en puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
