import time
from adapter import ChannelAdapter

adapter = ChannelAdapter()

# Simular envío de mensajes
eventos = [
    ("nueva_orden", {"id": 101, "cliente": "Juan", "monto": 50}),
    ("notificacion", {"mensaje": "Orden completada", "destino": "Juan"}),
    ("promocion", {"mensaje": "Hoy en tus compras 2x1", "destino": "Todos"})
]

for evento, data in eventos:
    adapter.publish(evento, data)
    time.sleep(1)
