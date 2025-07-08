# Go Codebase

This document outlines the project layout for the **januaruwanda Go Codebase** service, a Go-based application designed for large-scale data processing and computer vision tasks.

# Architecture
![architecture](docs/architecture/architecture.svg)

- **MinIO**: An object storage service used for storing and retrieving large amounts of unstructured data, such as video files and analytics results.

- **Database**: The persistent storage layer for structured data, including user information, configuration settings, and analytics metadata.

# Directory Structure

```plaintext
.
├── Dockerfile                # Docker configuration file
├── Makefile                  # Makefile for build automation
├── README.md                 # Project documentation
├── cmd                       # Main application directory
│   └── app
│       └── main.go           # Entry point of the application
├── config.json               # Configuration file
├── deployments               # Deployment configuration
│   └── docker-compose.yml    # Docker Compose file
├── devcontainer.json         # Development container configuration
├── docs                      # Documentation files
│   ├── api
│   │   └── spec.yml          # API specification
│   └── architecture
│       └── architecture.drawio # Architecture diagram
├── go.mod                    # Go module file
├── go.sum                    # Go module dependencies
├── internal                  # Internal application code
│   ├── delivery              # HTTP delivery layer
│   │   └── http
│   │       ├── healthcheck_controller.go # Health check controller
│   │       ├── route
│   │       │   └── route.go  # Route definitions
│   │       └── user_controller.go  # User controller
│   ├── domain                # Domain models
│   │   └── user.go           # User domain model
│   ├── presenters            # Presenters for HTTP responses
│   │   ├── healthcheck.go    # Health check response presenter
│   │   ├── status.go         # Status response presenter
│   │   └── user.go           # User response presenter
│   ├── repositories          # Repository layer for data access
│   │   └── user
│   │       └── user.go       # User repository
│   ├── usecases              # Use case implementations
│   │   └── user
│   │       └── user.go       # User use case
│   └── utils                 # Utility functions
│       └── crypto.go         # Cryptographic utilities
├── migrations                # Database migration files
│   └── 20241010_init.sql     # Initial database migration
└── pkg                       # Package directory for shared code
    └── database
        └── connection.go     # Database connection logic

```


## Project Layout Overview

The following directories are essential to the **bigvision-operation** project:


### Root Directory

- **`Dockerfile`**: 
  - This file serves as the blueprint for building the Docker image of the application. It defines how the application and its dependencies are packaged and executed within a containerized environment. By using this Dockerfile, developers can ensure consistency across various deployment stages and simplify the process of deploying the application in different environments.

- **`Makefile`**: 
  - The Makefile acts as an automation tool that streamlines various build processes. It contains predefined tasks that can be executed directly from the command line, such as building the application, running tests, or creating a Docker image. This simplifies the development workflow by providing a single interface for multiple commands, enhancing productivity and reducing the likelihood of human error during builds.

- **`config.json`**: 
  - This configuration file is utilized to define environment-specific settings and parameters essential for the application's operation. By externalizing configuration, the application can adapt to different environments without requiring code changes.

- **`deployments/`**: 
  - This directory contains deployment configuration files that facilitate the deployment of the application. 
  - **`docker-compose.yml`**: 
    - A crucial file that defines services, networks, and volumes for multi-container Docker applications.

- **`devcontainer.json`**: 
  - This configuration file specifies the setup for development containers, particularly for developers using Visual Studio Code. It allows for a consistent development environment across different machines, ensuring that all developers work with the same configurations, extensions, and tools.

- **`docs/`**: 
  - This directory houses documentation files relevant to the project.
  - **`api/`**: 
    - Contains the API specifications that outline how the API is structured and how to interact with it. 
  - **`architecture/`**: 
    - This folder contains diagrams that illustrate the architectural design of the application.
    - **`architecture.drawio`**: 
      - A visual representation of the high-level architecture, which aids in understanding how various components of the application interact with one another.

- **`go.mod`**

- **`go.sum`**

## Application Structure

### `cmd/`
- **`app/`**: 
  - This directory contains the main application code.
  - **`main.go`**: 
    - The entry point of the application where the main logic begins execution. It typically initializes the application, sets up the necessary configurations, and starts the HTTP server or other core components.

### `internal/`
- **`delivery/`**: 
  - This directory serves as the HTTP delivery layer of the application, handling incoming requests and sending responses.
  - **`http/`**: 
    - Contains various HTTP-related components that process requests and responses.
    - **`healthcheck_controller.go`**: 
      - A controller dedicated to handling health check requests, ensuring the service is operational and responding as expected.
    - **`route/`**: 
      - Contains the definitions for routing incoming HTTP requests to the appropriate handlers.
      - **`route.go`**: 
        - This file defines the routing logic for the application, mapping URLs to specific controllers and actions.
    - **`user_controller.go`**: 
      - A controller that manages user-related requests, handling operations such as user registration, login, and data retrieval.

- **`domain/`**: 
  - This directory contains the domain models that represent the core business logic of the application.
  - **`user.go`**: 
    - Defines the user domain model, detailing the structure and behavior of user-related data, including attributes and methods associated with user entities.

- **`presenters/`**: 
  - This directory houses presenters that format responses for HTTP requests. Presenters help create a consistent API contract between the frontend and backend, ensuring that the data returned to the client is properly structured and easily consumable.

- **`repositories/`**: 
  - This directory represents the repository layer responsible for data access and interactions with the database.
  - **`user/`**: 
    - Contains the user repository logic.
    - **`user.go`**: 
      - This file defines methods for accessing, querying, and manipulating user data in the database, encapsulating all database operations related to users.

- **`usecases/`**: 
  - This directory contains implementations of various use cases, encapsulating the application's core business logic.
  - **`user/`**: 
    - Contains user use case logic that defines the operations and workflows associated with user management.
    - **`user.go`**: 
      - This file outlines the use case methods related to user operations, such as creating, updating, and retrieving user information.

- **`utils/`**: 
  - This directory is dedicated to utility functions that support various parts of the application, providing common functionalities that can be reused across different components.
  - **`crypto.go`**: 
    - Contains cryptographic utility functions that are essential for secure data handling, including methods for hashing, encryption, and decryption.

### `migrations/`
- **`20241010_init.sql`**: 
  - This SQL file contains the initial database migration script, defining the schema changes and initial data setup necessary for the database. Migrations help manage database versions and ensure that the database schema is in sync with the application's requirements.

## Package Directory

### `pkg/`
- **`database/`**: 
  - This directory contains shared code related to database interactions, facilitating data access and management.
  - **`connection.go`**: 
    - This file defines the logic for establishing and managing database connections, ensuring efficient and reliable access to the database throughout the application. Proper connection management is crucial for maintaining performance and scalability in data-driven applications.