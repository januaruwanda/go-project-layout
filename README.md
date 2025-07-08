# Bigsocial Codebase

# Architecture
![architecture](docs/assets/arch.png)
- **FE Dashboard**: The front-end interface for users to interact with the system, providing visualizations and controls for video analytics.

- **Bigsocial-codebase**: The core backend service that handles all video analytics and API analytics, ensuring efficient processing and management of video data.

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
