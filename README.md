# inscripcion_crud
API de inscripciones, Integración con CI
inscripcion_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/inscripcion_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `API_PORT`: Puerto asignado para la ejecución del API
 - `INSCRIPCION_CRUD__PGUSER`: Usuario de la base de datos
 - `INSCRIPCION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `INSCRIPCION_CRUD__PGURLS`: Host de conexión
 - `INSCRIPCION_CRUD__PGDB`: Nombre de la base de datos
 - `INSCRIPCION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
API_PORT=8095 INSCRIPCION_CRUD__PGUSER=postgres INSCRIPCION_CRUD__PGPASS=password INSCRIPCION_CRUD__PGURLS=localhost INSCRIPCION_CRUD__PGDB=local INSCRIPCION_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/inscripcion_crud/blob/dev/modelo_inscripcion_crud.png).
