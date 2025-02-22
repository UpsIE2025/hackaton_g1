# Caso de mensajeria Channel Adapter

## Historia de Usuario
**COMO** usuario del sistema,  
**QUIERO** un adaptador de mensajería que envíe notificaciones vía SMS, EMAIL y PUSH,  
**PARA** que las alertas lleguen al canal adecuado según el tipo de notificación.

## Criterios de Aceptacion
**Dado** que un cliente realiza una compra en la tienda online,  
**CUANDO** el sistema genera una notificación de "Pedido Confirmado"  
**ENTONCES** el adaptador debe enviar un email con los detalles del pedido al cliente y un SMS si el cliente optó por recibir alertas móviles.

### Ejemplo:
> Una empresa de comercio electrónico necesita enviar notificaciones automáticas a sus clientes sobre el estado de sus pedidos. Dependiendo del tipo de notificación, el mensaje debe enviarse por SMS, EMAIL o PUSH.

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en máquina local:

### 1. Iniciar el consumidor de notificaciones en una terminal, para poder el seguimiento de las solicitudes:
```
py suscribe.py
```
### 2. Iniciar el API con FASTAPI por el puerto 8080:
```
uvicorn publisherapi:app --host 0.0.0.0 --port 8000 --reload
```
### 3. Realizar la peticion CURL con una notificacion, definiendo el tipo y el mensaje:
```
curl --location 'http://localhost:8000/publish/' \
--header 'Content-Type: application/json' \
--data '{
    "event_type": "PUSH",
    "data": {"mensaje": "Su Pedido ha sido procesado", "destino": "CLIENTE0001"}
}'
```
