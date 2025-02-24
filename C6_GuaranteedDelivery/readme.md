# Caso de mensajeria GuaranteedDelivery

## Historia de Usuario
**COMO** usuario del sistema,  
**QUIERO** que los clientes puedan generar pedidos en cualquier momento del día,  
**PARA** garantizar disponibilidad y mejorar la experiencia del cliente.

## Criterios de Aceptacion
**DADO** que un cliente desea realizar un pedido,  
**CUANDO** ingresa al sistema y genera un nuevo pedido,
**ENTONCES** el pedido debe ser registrado en la base de datos y quedar disponible para su procesamiento sin importar la hora del día.

### Ejemplo:
Un restaurante permite a sus clientes realizar pedidos en cualquier momento del día a través de su aplicación móvil. El sistema debe recibir, almacenar y procesar los pedidos sin importar la hora, asegurando que los pedidos hechos fuera del horario de cocina sean gestionados correctamente para el siguiente turno..

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en máquina local:

### 1. Iniciar el API con FASTAPI por el puerto 8080:
```
uvicorn publisherapi:app --host 0.0.0.0 --port 8000 --reload
```
### 2. Realizar la peticion CURL con una notificacion, definiendo el tipo y el mensaje:
```
curl --location 'http://localhost:8000/publish/' \
--header 'Content-Type: application/json' \
--data '{
    "channel": "orders",
    "message": "Pedido: 1 combo alitas"
}'
```
```
curl --location 'http://localhost:8000/publish/' \
--header 'Content-Type: application/json' \
--data '{
    "channel": "orders",
    "message": "Pedido: 1 cajita feliz"
}'
```
### 3. Iniciar el consumidor de notificaciones en una terminal, para poder ver el procesamiento de solicitudes acumuladas:
```
py consumer.py
```
