## 1. Problemas de diseño
Has sido encargado de plantear una solución al siguiente problema.
Imagina que tenemos un sistema escrito en cualquier lenguaje de programación compilado,
este sistema se conecta a una base de datos SQL, el sistema funciona muy bien cuando hay
baja demanda de transacciones, pero cuando la cantidad de transacciones aumenta el sistema
deja de responder solicitudes, se encontró que la base de datos es el cuello de botella y no
acepta solicitudes en paralelo para completar las transacciones. ¿Qué solución plantearías
para receptar más cantidad de transacciones adaptándonos al cuello de botella de la base de
datos?


### Respuesta
Primero revisar que los indices se encuentren bien realizados, en caso de no tenerlos empezar a aplicarlos, ya que se optimizarían las consultas. 
Se podría particionar la base de datos pero llevaría mucho esfuerzo por lo que consideraría realizar un sistema de colas a través de querys como 
<pre><code>SELECT ... FOR UPDATE SKIP LOCKED </code></pre>
con el fin de bloquear las filas de una tabla mientras se esta en una transacción mientras se seleccionan 
las que no estén bloqueadas. 
Claramente lo mejor es realizar un análisis profundo de la base, para poder seleccionar las mejores estrategias a implementar tanto dentro de la base
como en el codigo de la app.
