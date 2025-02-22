from flask import Flask, request, jsonify
from confluent_kafka import Producer
import json

app = Flask(__name__)

# Configuración del productor de Kafka
conf = {
    'bootstrap.servers': 'localhost:9092',  # Kafka desde Docker (localhost)
    'client.id': 'inventory-producer' #Topico
}

producer = Producer(conf)

# Función de entrega de mensajes
def delivery_report(err, msg):
    if err is not None:
        print(f"Error al enviar mensaje: {err}")
    else:
        print(f"Mensaje entregado a {msg.topic()} [{msg.partition()}]")

# Endpoint para recibir los mensajes y enviarlos a Kafka
@app.route('/send_notification', methods=['POST'])
def send_notification():
    data = request.get_json()  # Obtener los datos JSON de la petición POST

    product_id = data.get('product_id')
    status = data.get('status')

    if not product_id or not status:
        return jsonify({'error': 'Faltan datos: product_id y status'}), 400

    message = {
        'product_id': product_id,
        'status': status,
        'timestamp': '2025-02-22T12:00:00'  # Ejemplo de timestamp
    }

    producer.produce('inventory-updates', key=str(product_id), value=json.dumps(message), callback=delivery_report)
    producer.flush()

    return jsonify({'message': 'Notificación enviada correctamente'}), 200

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
