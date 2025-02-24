import redis
import json
from storage import MessageStorage

class ReliableConsumer:
    def __init__(self, redis_host="localhost", redis_port=6379, channel="orders"):
        self.redis_client = redis.StrictRedis(host=redis_host, port=redis_port, decode_responses=True)
        self.channel = channel
        self.storage = MessageStorage()

    def process_message(self, message_id, message):
        """ Procesa el mensaje y confirma su entrega """
        data = json.loads(message)
        print(f"✅ Procesando mensaje: {data['message']}")

        # Marcar como entregado en SQLite
        self.storage.mark_as_delivered(message_id)

    def handle_pending_messages(self):
        """ Reintenta enviar mensajes no entregados después de una falla """
        print("🔄 Revisando mensajes no entregados...")
        for message_id, channel, message in self.storage.get_undelivered_messages():
            print(f"🔄 Reintentando mensaje: {message}")
            self.redis_client.publish(channel, message)
            self.process_message(message_id, message)

    def start_listening(self):
        """ Escucha mensajes nuevos en Redis y los procesa """
        pubsub = self.redis_client.pubsub()
        pubsub.subscribe(self.channel)

        print(f"📡 Escuchando en canal '{self.channel}'...")
        self.handle_pending_messages()  # Reenviar mensajes no entregados

        for message in pubsub.listen():
            if message["type"] == "message":
                message_id = None
                for stored_message in self.storage.get_undelivered_messages():
                    if json.loads(stored_message[2]) == json.loads(message["data"]):
                        message_id = stored_message[0]
                        break
                
                self.process_message(message_id, message["data"])

if __name__ == "__main__":
    consumer = ReliableConsumer()
    consumer.start_listening()
