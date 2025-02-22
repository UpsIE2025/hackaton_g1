from confluent_kafka import Producer
import json

# Configuración del productor
conf = {
    'bootstrap.servers': 'localhost:9092',  # Dirección de tu Kafka desde Docker (localhost)
    'client.id': 'inventory-producer'
}

producer = Producer(conf)

# Función de entrega de mensajes
def delivery_report(err, msg):
    if err is not None:
        print(f"Error al enviar mensaje: {err}")
    else:
        print(f"Mensaje entregado a {msg.topic()} [{msg.partition()}]")

# Enviar un mensaje
def send_inventory_update(product_id, status):
    message = {
        'product_id': product_id,
        'status': status,
        'timestamp': '2025-02-22T12:00:00'  # Ejemplo de timestamp
    }
    producer.produce('inventory-updates', key=str(product_id), value=json.dumps(message), callback=delivery_report)
    producer.flush()

# Ejemplo de uso
if __name__ == "__main__":
    send_inventory_update(1, "en stock")
    send_inventory_update(2, "agotado")
