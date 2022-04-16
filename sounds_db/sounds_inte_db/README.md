# Instructions
Puerto TCP a usar: 3306.
1. Crear la imagen Docker, ejecutando dentro del mismo directorio el siguiente comando: `docker build -t sounds_inte_db .`

2. Desplegar la base de datos, mediante el siguiente comando:
`docker run -d -t -i -p 3301:3301 --name sounds_inte_db sounds_inte_db`

3. Desplegar el cliente web de MySQL PhpMyAdmin, mediante el siguiente comando:
`docker run --name db_inte -d --link sounds_inte_db:db -p 8081:80 phpmyadmin`

4. Acceder a la base de datos, usando el cliente PhpMyAdmin, mediante el navegador
web: http://localhost:8081.

5. Iniciar sesión usando las credenciales definidas en el Dockerfile de la imagen de la base de datos.

6. Crear una tabla de prueba y realizar inserciones y consultas a la base de datos.
