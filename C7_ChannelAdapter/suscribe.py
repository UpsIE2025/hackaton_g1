from adapter import ChannelAdapter
from logic import EVENT_HANDLERS

adapter = ChannelAdapter()

def handle_message(message):
    """ Callback para procesar mensajes recibidos """
    event_type = message["event"]
    data = message["data"]

    if event_type in EVENT_HANDLERS:
        EVENT_HANDLERS[event_type](data)
    else:
        print(f"⚠ Evento desconocido: {event_type}")

# Suscribirse al canal y procesar eventos
adapter.subscribe(handle_message)
