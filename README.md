# organizacion_crud
API de organizacion, Integración con CI
organizacion_crud master/develop
 ## Requirements
Go version >= 1.9.
 ## Preparation:  
    Para usar el API, usar el comando:
        - go get github.com/udistrital/organizacion_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `ORGANIZACION_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `ORGANIZACION_CRUD__PGUSER`: Usuario de la base de datos
 - `ORGANIZACION_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `ORGANIZACION_CRUD__PGURLS`: Host de conexión
 - `ORGANIZACION_CRUD__PGDB`: Nombre de la base de datos
 - `ORGANIZACION_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
ORGANIZACION_HTTP_PORT=8095 ORGANIZACION_CRUD__PGUSER=postgres ORGANIZACION_CRUD__PGPASS=password ORGANIZACION_CRUD__PGURLS=localhost ORGANIZACION_CRUD__PGDB=local ORGANIZACION_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/organizacion_crud/blob/develop/modelo_organizacion_crud.PNG).
