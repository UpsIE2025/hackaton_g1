def enviar_push(data):
    """ Simula el envío de una notificación PUSH"""
    print(f"Enviando PUSH: {data}")

def enviar_sms(data):
    """ Simula el envío de una notificación SMS"""
    print(f"Enviando SMS: {data}")

def enviar_email(data):
    """ Simula el envío de una notificación EMAIL"""
    print(f"Enviando EMAIL: {data}")

# Diccionario que mapea eventos a funciones
EVENT_HANDLERS = {
    "PUSH": enviar_push,
    "SMS": enviar_sms,
    "EMAIL": enviar_email
}
