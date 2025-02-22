from fastapi import FastAPI
from pydantic import BaseModel
from adapter import ChannelAdapter

app = FastAPI()
adapter = ChannelAdapter()

# Modelo de datos para validar solicitudes
class Event(BaseModel):
    event_type: str
    data: dict

@app.post("/publish/")
def publish_event(event: Event):
    """ Publica un evento en el canal de mensajería """
    adapter.publish(event.event_type, event.data)
    return {"message": "Evento publicado correctamente", "event": event.event_type}

