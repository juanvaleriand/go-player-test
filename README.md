
## Project Overview

The "Go Player Test" project is organized into several directories to maintain a clean and modular structure. Here's a brief description of each directory:

### `cmd/`

This directory contains the main entry point of the application, where the `main.go` file resides. It is responsible for initializing the application and starting its execution.

### `internal/`

The `internal/` directory holds the core implementation of the application. It includes subdirectories for various components:

- `api/`: Contains modules related to the API layer, including handlers, models, repository, requests, routes, and services.

  - `handlers/`: Implements request handlers for different API endpoints.
  - `models/`: Defines data models used throughout the application.
  - `repository/`: Contains database interaction logic using the repository pattern.
  - `requests/`: Defines request structures for various API endpoints.
  - `routes/`: Specifies the routing configuration for the API.
  - `services/`: Implements business logic and acts as an intermediary between handlers and repositories.

- `config/`: Holds configuration files and settings for the application.

- `database/`: Manages database-related functionality.

- `utils/`: Houses utility functions and helper modules used across the application.

This project structure is designed to promote modularity, maintainability, and separation of concerns, making it easier to develop, test, and scale the application.


To run this project, you will need to add the following environment variables to your .env file.

```
# Database Configuration
DB_HOST=your_db_host
DB_PORT=your_db_port
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

# Redis Configuration
REDIS_ADDR=localhost:6379
REDIS_PASS=

# GIN Mode (debug or release)
GIN_MODE=debug

SECRET_KEY=eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTcwMDEzODg3MCwiaWF0IjoxNzAwMTM4ODcwfQ.26yZhAdAqXRvGXR9j_OMBtSzTrkBHEIfXX7OlTsTIgU
```

## Run Project

Clone the project.

```bash
  git clone https://github.com/juanvaleriand/go-player-test.git
```

Go to the project directory.

```bash
  cd go-player-test
```

Install GO packages.

```bash
  go get github.com/juanvaleriand/go-player-test
```

Start the server.

```bash
  go run cmd/main.go
```

## API Documentation

For detailed documentation on API endpoints, please refer to the [Postman Collection](https://documenter.getpostman.com/view/7215921/2s9YXpWKHR).