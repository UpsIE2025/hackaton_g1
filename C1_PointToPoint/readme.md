# Proyecto NombreDelProyecto

## Historia de Usuario
**Como** un vehiculo que tiene varios lectores tags,  
**quiero** poder enviar un mensaje a través de un canal punto a punto,
**para** que solo el servidor que autoriza salidas de un lector en especifico se autorice y debite saldo. 
 
### Ejemplo:
> Soy un vehiculo que tiene un tag pero este puede ser leido por varios lectores a la salida, y este tag tiene que unicamente consumir un determinado servicio para el debito de su saldo. Entonces solo se deberia consumir esta consulta en un canal punto a punto. 

## Respuesta esperada
Respuesta esperada:
El mensaje es enviado a través del canal punto a punto.
Solo el receptor autorizado consume el mensaje.
Los demás sistemas no reciben ni consumen el mensaje, lo que garantiza que solo un receptor procese el acceso.

## Como se ejecuta el proyecto



## POSTMAN: Peticiones

