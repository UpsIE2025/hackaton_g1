# Caso de mensajeria Point to Point

## Historia de Usuario
**Como** dueño de un vehiculo que desea ser monitorizado,  
**quiero** poder leer y procesar los sensores de mi vechiulo,  
**para** que se puede registrar cada dato y relice su analisis con algun receptor que procese un determinado tipo. 

### Ejemplo:
> Un vehiculo sale de viaje, este necesita que se vea su estado de gasolina, velocidad, estado de encendido y puertas que van transcurriendo segun el tiempo de viaje. Cada sensor podria levantar altertas, mensajes y analitica diferente segun el sensoor (tipo de dato) que este recibiendo.

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en tu máquina local:

### 1. Para ejecutar el proyecto, levantarle el docker compose y el receptor en un terminal 1
  # TERMINAL 1
 ```sh  
 docker compose up -d
 go run .\C3_DataType\receptor\main.go
 ```
 ### 2. para enviar el productor de eventos ejecutar en un terminal 2: 
  # TERMINAL 2
 ```sh  
go run .\C3_DataType\producer\main.go
```
 ### 3. Par ausar desde POSTMAN, enviar al topico de velocidad que diferencia entre el tipo de dato. 

## **POSTMAN**

### **Enviar un Mensaje (POST)**
**POST** `http://localhost:8080/send`
- **Body (JSON):**
```json
{
  "topic": "c3_speed-topic",
  "message": "80"
}
```
- **Respuesta esperada:**
```json
"Velocidad recibida: 80 km/h"
```
