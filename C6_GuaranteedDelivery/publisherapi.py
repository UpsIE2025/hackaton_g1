from fastapi import FastAPI
from pydantic import BaseModel
import redis
import json
from storage import MessageStorage

app = FastAPI()
storage = MessageStorage()
redis_client = redis.StrictRedis(host="localhost", port=6379, decode_responses=True)

class MessageRequest(BaseModel):
    channel: str
    message: str

@app.post("/publish/")
def publish_message(request: MessageRequest):
    """ Publica un mensaje en Redis y lo almacena en SQLite """
    message_json = json.dumps({"message": request.message})
    
    # Guardar en SQLite antes de enviar
    message_id = storage.save_message(request.channel, message_json)
    
    # Publicar en Redis
    redis_client.publish(request.channel, message_json)
    print(f"📤 Mensaje enviado a {request.channel}: {message_json}")

    return {"message": "Mensaje publicado correctamente", "message_id": message_id}
