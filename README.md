#### A Survey Application

This application allows users to create surveys, submit responses, and view results. It uses Redis for data storage and is designed to be simple and easy to use.

## Features

- Create surveys with multiple questions and options
- Submit responses to surveys
- View survey results in real-time
  
## Technologies Used
  
- **Golang**: The main programming language used for the backend.
- **Redis**: Used for storing survey data and responses.
  
## To Run the Application
  

-  Make sure you have Golang and Redis installed on your machine.
-  Clone the repository and navigate to the project directory.
-  Run the application using the command:
  
  ```bash
  docker-compose up --build
  ```
  
  OR for development purposes, you can start only the Redis service with:
  
  ```bash
  docker-compose up -d redis
  ```

- The application will be available at `http://localhost:8080`.
- You can use redis-cli to ping the Redis server.
  
  ```bash
  redis-cli ping
  ```

- POST requests can be made to create surveys and submit responses.

```bash
     curl -X POST http://localhost:8080/submit \
       -d "language=Go"
```

![Survey Application Demo](/static/preview.png)




