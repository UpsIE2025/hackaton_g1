# Caso de mensajeria Point to Point

## Historia de Usuario
**Como** líder de un grupo de desarrollo para una empresa de encomiendas,  
**quiero** enviar las encomiendas a otro sistema con la seguridad de que todas lleguen al mismo,  
**para** que todas las encomiendas sean procesadas.

## Instrucciones para Ejecutar el Servicio

Sigue los siguientes pasos para ejecutar el servicio en tu máquina local:

```sh 
# levanta el receptor de encomiendas
go run C5_DeadLetter/receiver/main.go 

# este servicio envía las encomiendas al receptor, y si el receptor
# no está disponible las almacena en kafka para para ser procesadas
# cuando el receptor de encomiendas está activo de nuevo.
go run C5_DeadLetter/sender/main.go
```

### Respuesta Esperada:

Cuando el emisor de encomiendas, envía una encomienda pero el receptor no está disponible, está encomienda se debe guardar en una cola y ser reenviada cuando el receptor vuelva a estar disponible.
