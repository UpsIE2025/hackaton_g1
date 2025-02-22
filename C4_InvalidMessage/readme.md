**Historia de Usuario: InvalidMessage**

**Como** desarrollador en una empresa de logística,  
**quiero** detectar y gestionar mensajes inválidos en el sistema de mensajería basado en Redis,  
**para** evitar que pedidos con datos incompletos o erróneos afecten la operación y asegurar un procesamiento eficiente de la información.  

### **Ejemplo**  
Carlos trabaja en una empresa de logística que gestiona miles de pedidos diarios a través de un sistema de mensajería basado en Redis. Sin embargo, algunos mensajes llegan con información incompleta o incorrecta, lo que provoca retrasos y errores en la asignación de pedidos.  

Para solucionar este problema, se desarrollará una funcionalidad que detecte mensajes inválidos, los almacene en Redis para su posterior análisis y notifique al equipo correspondiente. Esto permitirá que los errores se manejen de manera más eficiente, reduciendo el impacto en la operación y mejorando la calidad del servicio.