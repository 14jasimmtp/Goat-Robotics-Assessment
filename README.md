# Time Tracking API

This is a Time Tracking API built using Golang with a clean architecture approach. The project includes user authentication, project and task management, and time tracking functionalities. It uses Gin for HTTP routing, GORM for ORM, and PostgreSQL as the database.

## Features

- User authentication with JWT
- Project and task management
- Time tracking for projects and tasks
- Unit tests for API endpoints, database interactions, and business logic
- Dockerized setup for easy deployment

## Running the Project with Docker Compose

### Prerequisites

- Docker
- Docker Compose

### Steps

1. Clone the repository:
    ```bash
    git clone https://github.com/14jasimmtp/Goat-Robotics-Assessment.git
    cd Goat-Robotics-Assessment
    ```

2. Build and run the Docker containers:
    ```bash
    docker-compose up --build
    ```

3. The API will be available at `http://localhost:3000` and the PostgreSQL database at `localhost:5432`.


