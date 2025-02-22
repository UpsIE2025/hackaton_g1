# Caso de mensajeria Point to Point

## Historia de Usuario
**Como** dueño de un vehiculo que desea ser monitorizado,  
**quiero** poder leer y procesar los sensores de mi vechiulo,  
**para** que se puede registrar cada dato y relice su analisis con algun receptor que procese un determinado tipo. 

### Ejemplo:
> Un vehiculo sale de viaje, este necesita que se vea su estado de gasolina, velocidad, estado de encendido y puertas que van transcurriendo segun el tiempo de viaje. Cada sensor podria levantar altertas, mensajes y analitica diferente segun el sensoor (tipo de dato) que este recibiendo.

Se van a enviar

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en tu máquina local:

### 1. Para ejecutar el proyecto, darle
 # Para enviar mensajes a KAFKA desde el PRODUCER a los 4 topicos de ejemplos
 # Se usa el proucer para ejemplo, luego se agregara para que este reciba datos desde el postman  
go run .\C3_DataType\producer\main.go

### Respuesta Esperada:
