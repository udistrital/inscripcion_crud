# admisiones_crud
API de admisiones, Integración con CI
admisiones_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/admisiones_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `ADMISIONES_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `ADMISIONES_CRUD__PGUSER`: Usuario de la base de datos
 - `ADMISIONES_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `ADMISIONES_CRUD__PGURLS`: Host de conexión
 - `ADMISIONES_CRUD__PGDB`: Nombre de la base de datos
 - `ADMISIONES_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
ADMISIONES_HTTP_PORT=8095 ADMISIONES_CRUD__PGUSER=postgres ADMISIONES_CRUD__PGPASS=password ADMISIONES_CRUD__PGURLS=localhost ADMISIONES_CRUD__PGDB=local ADMISIONES_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/admisiones_crud/blob/dev/modelo_admisiones_crud.png).
