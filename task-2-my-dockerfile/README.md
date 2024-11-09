# Build Dockerfile ke Docker Image

### Create Golang Project
---------------------------------------------------------------------------
![Screenshot Golang Project](img/1.png)
When run the Go project, open the browser. This will show in the browser
![Screenshot Golang Project](img/2.png)

### Add AUTHORS.md
Create AUTHORS.md. the following file will copy to docker image
![Screenshot AUTHORS.md](img/3.png)

### Build the Dockerfile
The Dockerfile is created to build the Docker image
![Screenshot Dockerfile](img/4.1.png)

### Build Docker Image
execute it using : docker build -t my-wise-word
![Screenshot execute command](img/4.png)

### Show in the Image
When open in images in docker desktop, this will show up
![Screenshot Images](img/5.png)

### Run New Container
using command : docker run -dp 8080:80 task-2-my-dockerfile my-wise-word 
![Screenshot Images](img/6.png)

### See The Logs
This is the logs when already run the container and exposed the app
![Screenshot Images](img/7.png)
![Screenshot Images](img/8.png)
