<h1 align="center"> Cuadrangular de Futbol API </h1>
 ##Índice

   *[Descripción de la aplicación](#descripción-de-la-aplicacion)
   
   *[Tecnologías utilizadas](#tecnologías-utilizadas)
   
   *[Puesta en Marcha](#puesta_en_marcha)

   <h2>Descripción de la aplicación</h2>
Este repositorio contiene una API sencilla para registrar los equipos de un cuadrangular de futbol (Nombre, Puntos, Goles a favor, Goles en contra), retorna en orden descendiente según los puntos obtenidos y en segunda prioridad por goles a favor.

Características principales:

  1. Ruta /api/teams: Permite registrar los equipos por nombre y les asigna un ID que se usa para las demás operaciones en la BD.

  2. Ruta /api/matches: A este endpoint se envían los datos de los partidos (Equipo A VS Equipo B, marcador), se registran 2 partidos (4 Equipos) en cada consulta.

  3. Ruta /api/standings: Retorna la información de los Equipos por Puntaje en orden Descendente.

<h2>Tecnologías utilizadas</h2>

1. Base de datos PostgreSQL: La aplicación utiliza PostgreSQL como base de datos usando dos tablas o entidades:

   - matches: registra la información de los partidos (ID del Equipo, puntos, goles a favor en la contienda, goles en contra en la contienda), que se van sumando a la siguiente entidad.
   - teams: registra los datos de cada equipo de fútbol: ID, Nombre, Puntos (Clasificación general), Goles en contra(del campeonato), Goles en contra(Del campeonato).
2. Se utiliza el ORM GORM para interactuar con la base de datos y proporciona soporte para migración de la base de datos.

3. Implementación de API con Gin: La aplicación utiliza el framework web Gin para implementar la API RESTful. Gin proporciona una sintaxis sencilla y un rendimiento eficiente para manejar las rutas, los middleware y las respuestas JSON.

<h2>Puesta en Marcha</h2>

Para poner en marcha el proyecto, se siguen los siguientes pasos:

 1. Clonar el repositorio desde el repositorio de GitHub.
```
git clone <URL del repositorio>

```
2. Asegúrese de tener Go instalado en su máquina. Puede descargarlo desde el sitio web oficial de Go: https://golang.org/dl/
3. Configurar la base de datos PostgreSQL:
     a. Instalar PostgreSQL en su computadora si aún no lo ha hecho. Puede descargarlo desde el sitio web oficial de PostgreSQL: https://www.postgresql.org/download/
     b. Crear una base de datos en PostgreSQL con el nombre "cuadrangular".
     c. Abrir el archivo db/database.go y actualiza los valores de user y password en la cadena de conexión DSN de ConnectDB() con tus propias credenciales de PostgreSQL.
4. Abrir una terminal y navega hasta el directorio raíz del proyecto.
   ````
   cd <ruta-del-proyecto>

   ```
5. Instalar las dependencias del proyecto.
   ```
   go mod init
   go mod tidy
   ó
   go get -u github.com/gin-gonic/gin
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/mysql
   ```

6. Inicia el servidor.
   ```
   go run main.go
   ```
   El servidor se iniciará en el puerto 8080.

   Ahora se puede probar la API usando Postman en las siguientes rutas:

   - POST: localhost:8080/api/teams  USANDO un JSON como el siguiente:
     [
        {"Nombre": "Equipo A"},
        {"Nombre": "Equipo B"},
        {"Nombre": "Equipo C"},
        {"Nombre": "Equipo D"}
      ]

   -POST: localhost:8080/api/matches USANDO un JSON como el siguiente:

   [
    {
        "LocalTeamID": 1,
        "VisitanteTeamID": 2,
        "GolesDelLocal": 0,
        "GolesDelVisitante": 0
    },
    {
        "LocalTeamID": 3,
        "VisitanteTeamID": 4,
        "GolesDelLocal": 2,
        "GolesDelVisitante": 0
    }
]

- GET: localhost:8080/api/standings con esta consulta se obtiene un JSON con los equipos ordenados como se mencionó anteriormente.

Para ver la Aplicación completa dirigirse al repositorio que contiene la interfaz cliente https://github.com/Astaroth-Andres-Molano/cuadrangular_cliente_app.git

           
