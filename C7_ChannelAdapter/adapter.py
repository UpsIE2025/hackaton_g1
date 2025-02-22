import redis
import json

class ChannelAdapter:
    def __init__(self, channel_name="app_channel", redis_host="localhost", redis_port=6379):
        self.channel_name = channel_name
        self.redis_client = redis.StrictRedis(host=redis_host, port=redis_port, decode_responses=True)

    def publish(self, event_type, data):
        """ Publica un mensaje en el canal """
        message = json.dumps({"event": event_type, "data": data})
        self.redis_client.publish(self.channel_name, message)
        print(f"📤 Publicado en canal '{self.channel_name}': {message}")

    def subscribe(self, callback):
        """ Suscribe al canal y ejecuta un callback al recibir mensajes """
        pubsub = self.redis_client.pubsub()
        pubsub.subscribe(self.channel_name)

        print(f"📡 Escuchando en canal '{self.channel_name}'...")

        for message in pubsub.listen():
            if message["type"] == "message":
                data = json.loads(message["data"])
                callback(data)

