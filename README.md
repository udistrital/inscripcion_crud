# inscripcion_crud

--Api de inscripciones con CI--
CI deploy with lambda - S3
Drone 0.8 
admisiones_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/inscripcion_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `INSCRIPCION_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `INSCRIPCION_CRUD__PGUSER`: Usuario de la base de datos
 - `INSCRIPCION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `INSCRIPCION_CRUD__PGURLS`: Host de conexión
 - `INSCRIPCION_CRUD__PGDB`: Nombre de la base de datos
 - `INSCRIPCION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: INSCRIPCION_CRUD_HTTP_PORT=8083 INSCRIPCION_CRUD__PGUSER=user INSCRIPCION_CRUD__PGPASS=password INSCRIPCION_CRUD__PGURLS=localhost INSCRIPCION_CRUD__PGDB=udistrital_core_db INSCRIPCION_CRUD__SCHEMA=core_new bee run

## MODELO
![inscripcion](https://user-images.githubusercontent.com/14035745/61454725-4fd9c000-a927-11e9-9009-36701f500135.png)

