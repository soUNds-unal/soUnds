# sounds_inte
Sounds interaction component.
* Base de Datos sounds_inte_db (MySQL): 
  * `puerto 3301`
  *  Etiqueta imagen docker: `sounds_inte_db`
  *  Etiqueta phpmyadmin imagen docker: `sounds_inte_db_phpmyadmin`
* Microservicio sounds_inte_ms (JavaScript + API-REST): 
  * `puerto 3302`
  *  Etiqueta imagen docker: `sounds_inte_ms`

## Instructions
Puerto TCP a usar: 3301.
1. (Omitir paso si ya se creó la imagen —chequear con `docker images`) 

   Crear la imagen Docker, ejecutando dentro del mismo directorio el siguiente comando: 

        docker build -t sounds_inte_db .

2. (Omitir paso si ya se creó el container —chequear con `docker ps -a`) 
   
   Desplegar la base de datos, mediante el siguiente comando:

        docker run -d -t -i -p 3301:3301 --name sounds_inte_db sounds_inte_db

3. (Omitir paso si ya está corriendo container —chequear con `docker ps`) 

   Correr container:

       docker container start sounds_inte_db

3. Desplegar el cliente web de MySQL PhpMyAdmin, mediante el siguiente comando:

        docker run --name phpmyadmin -d --link sounds_inte_db:db -p 8081:80 phpmyadmin

4. Acceder a la base de datos, usando el cliente PhpMyAdmin, mediante el navegador
web: http://localhost:8081.

5. Iniciar sesión usando las credenciales definidas en el Dockerfile de la imagen de la base de datos.
  * `Username: sounds_inte`
  * `Password: 2022`

6. Crear una tabla de prueba y realizar inserciones y consultas a la base de datos.
