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

