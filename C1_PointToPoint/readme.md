# Proyecto NombreDelProyecto

## Historia de Usuario
**Como** líder de un grupo de desarrollo para una empresa de encomiendas,  
**quiero** enviar las encomiendas a un solo receptor,  
**para** para que una misma encomienda no se procese más de una vez por varios receptores.

## Respuesta esperada
Cada encomienda debe ser procesada por un solo receptor.
Cada encomienda solo debe enviarse a un solo receptor.

## Como se ejecuta el proyecto

```bash
# Levantar 2 receptores con
go run receiver/main.go 

# Ejecutar el emisor (el cual contiene un servidor http)
go run sender/main.go
```

Desde Postman enviar la siguiente solicitud:
`POST http://localhost:80/`
```json
{
    "id": 1,
    "from": "city a",
    "to": "city b",
    "size": "10 kg"
}
```

## POSTMAN: Peticiones

Ver archivo `C1_PointToPoint.postman_collection.json`
