# sounds_inte
Sounds interaction component.
* Base de Datos sounds_inte_db (MySQL): 
  *  Puerto: `3301`
  *  Etiqueta imagen docker: `sounds_inte_db`
* phpmyadmin (MySQL): 
  *  Puerto: `8081`
  *  Etiqueta imagen docker: `phpmyadmin`
* Microservicio sounds_inte_ms (JavaScript + API-REST): 
  *  Puerto: `3302`
  *  Etiqueta imagen docker: `sounds_inte_ms`


## INSTRUCCIONES DESPLEGAR sounds_inte_db Y phpmyadmin
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


## INSTRUCCIONES DESPLEGAR sounds_inte_ms

In order to dockerize, we follow the resource: https://www.digitalocean.com/community/tutorials/how-to-build-a-node-js-application-with-docker

Puerto TCP a usar: 3302.

1. (Omitir paso si ya se creó la imagen —chequear con `docker images`) 

   Crear la imagen Docker, ejecutando dentro del mismo directorio el siguiente comando: 

        docker build -t sounds_inte_ms .

2. (Omitir paso si ya se construyó el container —chequear con `docker ps -a`) 
   
   Construir (build) container:

        docker run --name sounds_inte_ms -p 3302:3302 -e DB_HOST=172.17.0.2 -e DB_PORT=3301 -e DB_USER=sounds_inte -e DB_PASSWORD=2022 -e DB_NAME=sounds_inte_db -e URL=0.0.0.0:3302 sounds_inte_ms

3. (Omitir paso si ya está corriendo container —chequear con `docker ps`) 

   Correr container:

       docker container start sounds_inte_ms

  Explicación:
     
      docker run -p 3302:3302 -e DB_HOST=X -e DB_PORT=3301 -e DB_USER=Y -e DB_PASSWORD=Z -e DB_NAME=sounds_inte_db -e URL=0.0.0.0:3302 sounds_inte_ms
      
      X = IP del contenedor (sounds_inte_db)
      X = docker inspect container_name_or_id
      
      Ver atributo Networks > bridge > IPAddress
      X = 172.17.0.2

      Y = usuario de la base de datos
      Y = sounds_inte

      Z = contraseña de la base de datos
      Z = 2022
