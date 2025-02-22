# Caso de mensajeria Point to Point

## Historia de Usuario
**Como** gerente de una tienda en línea,
**Quiero** que los pedidos realizados en nuestro sistema de e-commerce (que usa Kafka) se repliquen automáticamente en nuestro sistema de gestión de inventario (que usa Redis) mediante un Messaging Bridge,
**Para** asegurar que el stock se actualice en tiempo real sin intervención manual.

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en tu máquina local:

```sh
go run ecommerce/main.go
go run inventario/main.go
go run bridge/main.go
```

Luego se envía una solicitud de pedido al servicio de ecommerce:
`POST http://localhost:80/`
```json
{
    "id": 1,
    "products": [
        {
            "id": 201,
            "name": "pelota",
            "quantity": 3
        },
        {
            "id": 205,
            "name": "mesa",
            "quantity": 1
        }
    ]
}
```

El servicio de e-commerce enviará este pedido al topic `orders` usando `Kafka` el cual será leído y transformado por nuestro bridge para su envío al topic `descreaseStock` través de `Redis` al servicio de inventario.

### Respuesta Esperada:

Un pedido generado en el sistema de e-commerce debe producir un mensaje en Kafka, y ese mensaje debe ser transformado y reenviado correctamente al sistema de inventario (Redis) a través del Messaging Bridge.

## POSTMAN: Peticiones

Ver archivo `C8_MesagginBridge.postman_collection.json`
