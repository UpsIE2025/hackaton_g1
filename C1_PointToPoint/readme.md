# Proyecto NombreDelProyecto

## Historia de Usuario
**Como** usuario del sistema,  
**quiero** poder enviar un mensaje a través de un canal punto a punto,
**para** que solo el destinatario especificado reciba el mensaje y lo consuma de manera eficiente, sin interferencias de otros
 
### Ejemplo:
> Imagina que eres un operador de un sistema de mensajería dentro de un entorno de control vehicular. Necesitas enviar un mensaje de autorización para que un vehículo pueda acceder a una zona restringida, pero no quieres que otros operadores reciban y procesen el mensaje, solo el sistema de acceso adecuado. En este caso, tú envías el mensaje a través de un canal punto a punto, y el sistema de acceso correcto recibe el mensaje y autoriza el acceso, mientras que los demás sistemas no interactúan con este mensaje.

## Respuesta esperada
Respuesta esperada:
El mensaje es enviado a través del canal punto a punto.
Solo el receptor autorizado (en este caso, el sistema de acceso correspondiente) consume el mensaje.
Los demás sistemas no reciben ni consumen el mensaje, lo que garantiza que solo un receptor procese el acceso.


## Como se ejecuta el proyecto



## POSTMAN: Peticiones

