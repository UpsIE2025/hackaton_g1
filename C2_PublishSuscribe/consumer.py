from confluent_kafka import Consumer, KafkaException, KafkaError
import json

# Configuración del consumidor
conf = {
    'bootstrap.servers': 'localhost:9092',  # Kafka desde Docker (localhost)
    'group.id': 'inventory-consumers',  # Topico (grupo de consumidores)
    'auto.offset.reset': 'earliest'  
}

consumer = Consumer(conf)

# Suscribirse al topic
consumer.subscribe(['inventory-updates'])

# Función para procesar los mensajes
def process_message(message):
    data = json.loads(message.value().decode('utf-8'))
    print(f"Producto ID: {data['product_id']} - Estado: {data['status']}")

# Escuchar los mensajes
try:
    while True:
        msg = consumer.poll(timeout=1.0)  # Tiempo de espera de 1 segundo por mensaje
        if msg is None:
            continue
        if msg.error():
            if msg.error().code() == KafkaError._PARTITION_EOF:
                # Fin de la partición, no hacer nada
                continue
            else:
                raise KafkaException(msg.error())
        else:
            # Procesar el mensaje
            process_message(msg)

except KeyboardInterrupt:
    print("Interrumpido por el usuario.")
finally:
    # Cerrar el consumidor
    consumer.close()
