**Historia de Usuario: InvalidMessage**

**Como** desarrollador en una empresa de logística,  
**quiero** detectar y gestionar mensajes inválidos en el sistema de mensajería basado en Redis,  
**para** evitar que pedidos con datos incompletos o erróneos afecten la operación y asegurar un procesamiento eficiente de la información.  

### **Ejemplo**  
Carlos trabaja en una empresa de logística que gestiona miles de pedidos diarios a través de un sistema de mensajería basado en Redis. Sin embargo, algunos mensajes llegan con información incompleta o incorrecta, lo que provoca retrasos y errores en la asignación de pedidos.  

Para solucionar este problema, se desarrollará una funcionalidad que detecte mensajes inválidos, los almacene en Redis para su posterior análisis y notifique al equipo correspondiente. Esto permitirá que los errores se manejen de manera más eficiente, reduciendo el impacto en la operación y mejorando la calidad del servicio.

## **Funcionamiento del sistema**

El sistema desarrollado es una **aplicación Java independiente** que utiliza **Redis** para gestionar mensajes válidos e inválidos. Se implementa un **servidor HTTP embebido** en Java que expone endpoints REST para interactuar con Redis.

El flujo de trabajo del sistema es el siguiente:
1. **El usuario envía un mensaje** a través del endpoint REST.
2. **El sistema evalúa el mensaje** para determinar si es válido o inválido.
3. **Si el mensaje es inválido**, se almacena en la lista `invalid_messages` en Redis.
4. **Si el mensaje es válido**, se almacena en la lista `valid_messages` en Redis.
5. **El usuario puede consultar** los mensajes válidos e inválidos mediante los endpoints REST.

---

## **Instalación y ejecución**

### **1. Requisitos previos**
- Java 11 o superior
- Maven
- Docker (para ejecutar Redis)

### **2️. Levantar Redis con Docker**
```sh
 docker run --name redis-server -p 6379:6379 -d redis
```

### **3️. Construcción y ejecución del proyecto**
```sh
mvn clean package
mvn exec:java -Dexec.mainClass=RedisMessageProcessor
```

Si la ejecución es correcta, verás en la consola:
```sh
✅ Servidor REST en ejecución en http://localhost:8080/
```

---

## **POSTMAN**

### **Enviar un Mensaje (POST)**
**POST** `http://localhost:8080/messages/send`
- **Body (JSON):**
```json
{
    "id": 1,
    "sensor": "fuel",
    "value": 50
}
```
- **Respuesta esperada:**
```json
"✅ Mensaje válido almacenado en Redis."
```

