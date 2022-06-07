<<<<<<< HEAD
# Instructions
Puerto TCP a usar: 3301.
1. Crear la imagen Docker, ejecutando dentro del mismo directorio el siguiente comando:
`docker build -t swarch2022i_db .`
2. Desplegar la base de datos, mediante el siguiente comando:
`docker run -d -t -i -p 3306:3306 --name swarch2022i_db swarch2022i_db`
3. Desplegar el cliente web de MySQL PhpMyAdmin, mediante el siguiente comando:
`docker run --name db_client -d --link swarch2022i_db:db -p 8081:80 phpmyadmin`
4. Acceder a la base de datos, usando el cliente PhpMyAdmin, mediante el navegador
web: http://localhost:8081.
5. Iniciar sesión usando las credenciales definidas en el Dockerfile de la imagen de la base de datos.
6. Crear una tabla de prueba y realizar inserciones y consultas a la base de datos.
=======
# sounds_inte_ms
In order to dockerize, we follow the resource: https://www.digitalocean.com/community/tutorials/how-to-build-a-node-js-application-with-docker

1. Build docker image: `docker build -t sounds_inte_ms .`
2. Check images: `docker images`
3. Build the container: `docker run -p 8080:8080 -e DB_HOST=172.17.0.2 -e DB_PORT=3301 -e DB_USER=sounds_inte -e DB_PASSWORD=2022 -e DB_NAME=sounds_inte_db -e URL=0.0.0.0:8080 sounds_inte_ms`
  
  Explicación:
     
      docker run -p 2002:2002 -e DB_HOST=X -e DB_PORT=3301 -e DB_USER=Y -e DB_PASSWORD=Z -e DB_NAME=sounds_inte_db -e URL=0.0.0.0:2002 sounds_inte_ms
      
      X = IP del contenedor (sounds_inte_db)
      X = docker inspect container_name_or_id
      
      Ver atributo Networks > bridge > IPAddress
      X = 172.17.0.2

      Y = usuario de la base de datos
      Y = sounds_inte

      Z = contraseña de la base de datos
      Z = 2022

>>>>>>> 4de141db9792184083a0fb84c6088fbed0c060bb
