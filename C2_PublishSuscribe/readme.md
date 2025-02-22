# Caso de mensajeria Publish/Suscribe

## Historia de Usuario

**Como** responsable de la plataforma de notificaciones de una tienda en línea  
**quiero** implementar un sistema de notificación basado en el patrón Publish/Subscribe  
**para** que cuando haya actualizaciones en el inventario de productos, los usuarios interesados (como el personal de ventas, administración o clientes) reciban automáticamente las actualizaciones en tiempo real sin necesidad de hacer consultas repetitivas.


## Ejemplo:
Para mejorar la eficiencia y garantizar que las notificaciones sean inmediatas y en tiempo real, necesitamos implementar un **patrón de diseño Publish/Subscribe**. Este patrón permitirá que el sistema publique eventos de actualización de inventario, y que los usuarios interesados (empleados de ventas, administradores y clientes) reciban estas notificaciones de forma automática cuando se produzcan cambios en los productos que siguen o gestionan.

El sistema de inventario (Publisher) publicará eventos que se distribuirán a todos los suscriptores interesados en esos cambios (como los empleados de ventas, administradores o clientes que siguen productos específicos).

## Respuesta esperada:

Cuando se implemente el sistema de notificaciones, se tendrá:

- **Los suscriptores** recibirán mensajes automáticamente cuando se produzca una actualización en el inventario (por ejemplo, cuando un producto vuelva a estar disponible o se agote).
- **Los clientes** recibirán las notificaciones en tiempo real sobre los productos que están siguiendo.
- **Los empleados de ventas** estarán al tanto de los cambios de stock en tiempo real, lo que les permitirá ofrecer productos disponibles de manera más eficiente.
- **El administrador** podrá recibir actualizaciones sobre todos los productos sin necesidad de consultas activas.

El sistema aprovechará **Kafka** para garantizar que los eventos de inventario se distribuyan de forma eficiente a todos los suscriptores, sin sobrecargar el sistema con consultas repetitivas.

## Cómo se ejecuta la implementación:

1. **Kafka (Publisher y Broker)**:
   - **El sistema de inventario (Publisher)**: El sistema publicará eventos sobre los cambios de inventario en un **topic** específico de Kafka. Por ejemplo, cada vez que un producto cambie de estado, se publicará un mensaje en el topic `inventory-updates`.
   - **Suscriptores**: Los suscriptores (empleados de ventas, administradores y clientes) se suscribirán al topic correspondiente para recibir los eventos en tiempo real.

2. **Kafka Producer**:
   - El **Producer** es responsable de enviar los eventos de cambio de inventario a Kafka. Cada vez que un producto cambie de estado (por ejemplo, se vuelva disponible o se agote), el **Producer** enviará un mensaje al **topic** `inventory-updates`.
   - El mensaje incluirá detalles como el `product_id`, y `status` de la actualización.

3. **Kafka Consumer**:
   - Los **Consumers** (empleados de ventas, administradores, y clientes) se suscribirán a los eventos del topic `inventory-updates` para recibir las actualizaciones correspondientes.
   - Cada suscriptor recibirá los mensajes del topic y podrá procesarlos según su lógica (notificación a clientes, actualización de stock, etc.).

4. **Manejo de los eventos de inventario**:
   - El sistema publicador enviará eventos sobre los cambios en el inventario.
   - Los consumidores recibirán los eventos en tiempo real y realizarán las acciones necesarias, como notificar a los usuarios.

5. **Kafka Topic**:
   - El **topic** `inventory-updates` será utilizado para distribuir todos los eventos de inventario a los suscriptores.

6. **El archivo `app.py`**:
   - El archivo `app.py` es el servidor **Flask** que actúa como el **Producer** en el sistema. Se encarga de recibir solicitudes HTTP en el endpoint `/send_notification` para registrar los cambios en el inventario de productos. Una vez recibe la solicitud, envía los detalles del producto y su estado al **topic `inventory-updates`** de Kafka. Los **Consumers** suscritos recibirán estos mensajes en tiempo real.

## Cómo validar desde POSTMAN: Peticiones

Aunque **Postman** no es la herramienta ideal para interactuar directamente con **Kafka**, puedes crear endpoints RESTful para simular la interacción con el sistema de Kafka mediante HTTP.

### Validación con Kafka:

1 **Publicación de un evento de inventario (Publisher de Kafka)**:
   - El **Publisher** publicará el evento de inventario en Kafka mediante una API que se encargue de enviar el evento al **topic**.
   - **POST** `http://localhost:5000/send_notification`
     - Body:
     ```json
     {
       "product_id": "1",
       "status": "en stock"
     }
     ```
   Este endpoint simula la actualización de inventario y la publicación del evento al **topic** `inventory-updates` en Kafka.

### Validación en Kafka:

1. **Verificación de la publicación**:
   - Puedes usar herramientas como **Kafka-console-consumer** para verificar que los mensajes se estén publicando correctamente en el topic `inventory-updates`.
     ```bash
     kafka-console-consumer --bootstrap-server localhost:9092 --topic inventory-updates --from-beginning
     ```

2. **Verificación de los consumidores**:
   - Si has implementado consumidores en tu aplicación, puedes verificar que estén recibiendo los mensajes y procesándolos correctamente. Los consumidores deberían mostrar los mensajes recibidos en tiempo real.
