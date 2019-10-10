# organizacion_crud
API de gestión de organizaciones

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `organizacion_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/organizacion_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `ORGANIZACION_CRUD__HTTP_PORT`: Puerto asignado para la ejecución del API
 - `ORGANIZACION_CRUD__DB_USER`: Usuario de la base de datos
 - `ORGANIZACION_CRUD__DB_PASS`: Clave del usuario para la conexión a la base de datos  
 - `ORGANIZACION_CRUD__DB_URL`: Host de conexión
 - `ORGANIZACION_CRUD__DB_NAME`: Nombre de la base de datos
 - `ORGANIZACION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
ORGANIZACION_CRUD__HTTP_PORT=8097 ORGANIZACION_CRUD__DB_USER=user ORGANIZACION_CRUD__DB_PASS=password ORGANIZACION_CRUD__DB_URL=localhost ORGANIZACION_CRUD__DB_NAME=bd ORGANIZACION_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/organizacion_crud/blob/develop/modelo_organizacion_crud.png).
