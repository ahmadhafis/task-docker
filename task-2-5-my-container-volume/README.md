# Container, Env, & Volume

### Create new container with env and Volume
---------------------------------------------------------------------------
Run the following command to create a new container with env and a volume:
*docker run -d --name my-postgres-v2-ahmadhafis \
(name-of-container) \
-e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password \
(username & password for postgredb) \
-v my-pg-volume-ahmadhafis:/var/lib/postgresql/data \
(name of volume and where env will be mounted) \
-p 5434:5432 postgres (port local:docker)*
![Screenshot Golang Project](img/1.png)
Result
![Screenshot Golang Project](img/2.png)
![Screenshot Golang Project](img/3.png)

### Check the Postgres in DBeaver
result
![Screenshot AUTHORS.md](img/4.png)
Add new table and insert 10 records
![Screenshot AUTHORS.md](img/5.png)
![Screenshot AUTHORS.md](img/6.png)

### Delete the Container and Create new container
stop and Delete the container
![Screenshot Dockerfile](img/7.png)
Create new container with name *my-postgres-v2-ahmadhafis* with the sam env and volume
![Screenshot Dockerfile](img/8.png)

### See if the table still exist
![Screenshot execute command](img/9.png)
