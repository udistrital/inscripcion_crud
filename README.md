# admisiones_crud

--Api de personas con CI--
CI deploy with lambda - S3
Drone 0.8 
admisiones_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/admisiones_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_ADMISIONES_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `ADMISIONES_CRUD__PGUSER`: Usuario de la base de datos
 - `ADMISIONES_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `ADMISIONES_CRUD__PGURLS`: Host de conexión
 - `ADMISIONES_CRUD__PGDB`: Nombre de la base de datos
 - `ADMISIONES_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_ADMISIONES_HTTP_PORT=8083 ADMISIONES_CRUD__PGUSER=user ADMISIONES_CRUD__PGPASS=password ADMISIONES_CRUD__PGURLS=localhost ADMISIONES_CRUD__PGDB=udistrital_core_db ADMISIONES_CRUD__SCHEMA=core_new bee run

## MODELO
![image]().
