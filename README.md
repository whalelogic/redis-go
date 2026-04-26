A Survey Application demo built with Golang and Redis.

This application allows users to create surveys, submit responses, and view results. It uses Redis for data storage and is designed to be simple and easy to use.

## Features

- Create surveys with multiple questions and options
- Submit responses to surveys
- View survey results in real-time
  
  ## Technologies Used
  
- Golang: The main programming language used for the backend.
- Redis: Used for storing survey data and responses.
  
  ## To Run the Application
  

1. Make sure you have Golang and Redis installed on your machine.
2. Clone the repository and navigate to the project directory.
3. Run the application using the command:
  
  ```bash
  docker-compose up --build
  ```
  
  OR for development purposes, you can start only the Redis service with:
  
  ```bash
  docker-compose up -d redis
  ```
  
  > A better approach is to define env variables in a `.env` file and use them in the `docker-compose.yml` file. For example, you can create a `.env` file with the following content:
  
  ```env
  REDIS_ADDR=localhost:6379
  ```
  

- These values are _`hard-coded`_ into the `main.go` file. 
