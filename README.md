# inscripcion_crud
API de gestión de inscripciones y admisiones

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `inscripcion_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/inscripcion_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `INSCRIPCION_CRUD_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `INSCRIPCION_CRUD__PGUSER`: Usuario de la base de datos
 - `INSCRIPCION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `INSCRIPCION_CRUD__PGURLS`: Host de conexión
 - `INSCRIPCION_CRUD__PGDB`: Nombre de la base de datos
 - `INSCRIPCION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
INSCRIPCION_CRUD_HTTP_PORT=8887 INSCRIPCION_CRUD__PGUSER=user INSCRIPCION_CRUD__PGPASS=password INSCRIPCION_CRUD__PGURLS=localhost INSCRIPCION_CRUD__PGDB=bd INSCRIPCION_CRUD__SCHEMA=schema_new bee run

## MODELO
![image](https://github.com/planesticud/inscripcion_crud/blob/develop/modelo_inscripcion_crud.png).
