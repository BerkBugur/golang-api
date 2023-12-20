# Go-Project

This project is a RESTful API written in Go. It offers a simple and effective solution for task management. Users can create, update, delete and list tasks.

## Requirements

This section lists the requirements to run the project on your local machine

- You need to install [Docker](https://docs.docker.com/get-docker/) / [Docker Compose](https://docs.docker.com/compose/)

- Clone this source code your machine.

- If you are going to run the application over the cloud (e.g. AWS), **Allow Ports** in the security group so that the services can communicate with each other.

## While starting

In the **Project Directory**, you need to build first after that can run:

`docker compose build`

`docker compose up -d`

and wait some tine for build docker images  then 4 service will be start.
```bash
  [+] Running 4/4
 ✔ Container postgres                Started                               0.0s
 ✔ Container go-project-my-app-1     Started                               0.0s
 ✔ Container go-project-promethus-1  Started                               0.0s
 ✔ Container go-project-grafana-1    S...                                  0.0s

```
## API Documentation
Swagger 2.0 was used for documentation. You can view all APIs from the link below.

`http://<your-ip>:8080/swagger/index.html`

##@@IMAGE@@

## Testing
Lorem 


## Logging and Monitoring

### Monitoring

Grafana and Prometheus were used for monitoring. You can follow the instructions below to visualize and see metric data.

**Prometheus collects metrics**

``http://<your-ip>:8081/metrics ``

**Prometheus page**

``http://<your-ip>:9090/graph``


**Grafana login page**

*username: admin - pass: admin*

``http://<your-ip>:3000/login`` 

##@@IMAGE@@

### Logging
You can see the app.log file in the directory or we can connect to the container and see it. For this, we connect to the my-app container and then we can read the file.

```bash
docker exec -it go-project-my-app-1 /bin/sh

cat app.log
```

##@@IMAGE@@

## Project Security and JWT Usage

In this project, certain endpoints are secured using JSON Web Tokens (JWT). JWT usage provides protection against unauthorized access by users. Below is a step-by-step guide on how to securely access these endpoints:

- To test these processes, you can use tools such as Postman for managing API requests, or you can use the integrated Swagger interface in your project.
### Sign up Login and Test
- You need to register as a user in the system. To do this, send a POST request to the following endpoint:

  ``http://<your-ip>:8080/users/signup ``

- To log in to the system, send a POST request to the following endpoint:

  ``http://<your-ip>:8080/users/login ``

  > *Upon this request, you will be provided with a JWT token. This token can be stored as a **COOKIE**. If you encounter any issues with cookies, you can use the returned token for your tests.*

- To verify your login, send a GET request to the following endpoint to check the validity of the token:

  ``http://<your-ip>:8080/users/validate ``

### Example Usage

Once you complete the login process, you can access secure endpoints. If you do not log in, you will receive a "401 unauthorized" For example, to get a paginated task list:

``http://<your-ip>:8080/tasks/paged?page=1&size=10``

This request will return the first 10 tasks.

**For shut down this project go project directory**


``docker compose down``
