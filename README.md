# User and Role Management API (Clean Architecture)

Read this in [Bahasa Indonesia](README.id.md) 🇮🇩

## Introduction

Initially, this project was designed to implement **Role-Based Access Control (RBAC)** using JWT for authorization and route protection.

It has since evolved to adopt **Clean Architecture**, ensuring a more organized and standardized codebase. Additionally, the project is equipped with **Swagger** documentation for seamless API testing and integrates **Redis** for caching to optimize performance.

---

## Concepts & Architecture

This project rigorously adheres to **Clean Architecture** principles to achieve a clear **separation of concerns**, ensuring that each layer has a distinct responsibility. This approach keeps the codebase highly testable, maintainable, and scalable for future development.

The project structure is organized into the following layers:

1.  **Domain Layer**: The core of the application. This layer encapsulates **entities** (models) and **business logic** (services). It is completely **agnostic** to technical details, meaning it has no dependency on the database or external frameworks.
    * `internal/domain/{user,role,auth}/model.go`
    * `internal/domain/{user,role,auth}/service.go`
    * `internal/domain/{user,role}/repository.go` (Interface)

2.  **Infrastructure Layer**: Handles the technical implementation of the interfaces defined in the Domain Layer. This includes **database persistence** (PostgreSQL), **caching mechanisms** (Redis), and other external configurations.
    * `internal/infrastucture/repository/`
    * `internal/infrastucture/cache/`
    * `config/`
    * `utils/`

3.  **Presentation Layer**: The entry point for external interactions. In this project, this layer serves as the REST API handler, implemented using the **Gin Web Framework**.
    * `internal/domain/{user,role,auth}/handler.go`
    * `routes/routes.go`

This separation ensures that the core business logic remains **pure** and strictly **decoupled** from infrastructure concerns.

---

## Key Features

- **JWT-Based Authentication & Authorization**: A secure login system that generates JWT tokens for accessing protected endpoints.
- **User Management (Admin)**: Full CRUD operations to manage user data effectively.
- **Role Management (Admin)**: CRUD operations to manage roles and access privileges (e.g., `user` & `admin`).
- **Role-Based Access Control (RBAC)**: Middleware designed to restrict access to specific endpoints based on authorized roles.
- **Redis Caching Layer**: Implements caching on the user repository using the **Decorator Pattern**. This significantly reduces database load and improves response times, featuring smart **cache invalidation strategies**.
- **Robust Error Handling**: Utilizes **sentinel errors** to distinguish between business logic errors (e.g., _record not found_) and technical errors, ensuring accurate and meaningful HTTP response codes.
- **API Documentation (Swagger)**: Auto-generated, interactive API documentation for easy testing and integration.
- **Password Hashing**: Secure password storage using `bcrypt`.
- **Centralized Configuration**: Manages environment variables and configuration via `.env` files.

---


## API Documentation & Endpoints

Comprehensive API documentation is available via **Swagger**. Once the application is running, you can access the interactive documentation at:

### Ringkasan Endpoint

| Method | Endpoint | Description | Access |
| :--- | :--- | :--- | :--- |
| `POST` | `/register` | Register a new user | Public |
| `POST` | `/login` | Login & retrieve JWT token | Public |
| `GET` | `/user/profile` | Get current user profile | User (Authenticated) |
| `GET` | `/admin/users` | List all users | Admin |
| `POST` | `/admin/users` | Create a new user | Admin |
| `GET` | `/admin/users/{id}` | Get user by ID | Admin |
| `PUT` | `/admin/users/{id}` | Update user by ID | Admin |
| `DELETE`| `/admin/users/{id}`| Delete user by ID | Admin |
| `GET` | `/admin/roles` | List all roles | Admin |
| `POST` | `/admin/roles` | Create a new role | Admin |
| `GET` | `/admin/roles/{id}` | Get role by ID | Admin |
| `PUT` | `/admin/roles/{id}` | Update role by ID | Admin |
| `DELETE`| `/admin/roles/{id}`| Delete role by ID | Admin |

---

## Tech Stack

* **Language**: Golang
* **Framework**: Gin Gonic
* **Database**: PostgreSQL
* **Caching**: Redis
* **ORM**: GORM
* **Documentation**: Swaggo
* **Utilities**: `godotenv`, `jwt-go`, `bcrypt`
---

## Project Structure

The directory structure is meticulously designed to adhere to Clean Architecture principles, ensuring separation of concerns and maintainability.

```
user-role-management/
├── config/
│   ├── cache.go                        # Redis connection setup
│   ├── database.go                     # Database connection setup
│   └── env.go                          # Loads environment variables from .env
├── docs/
│   ├── docs.go                         # Main Swaggo generated file
│   ├── swagger.json                    # OpenAPI spec in JSON format
│   ├── swagger.yaml                    # OpenAPI spec in YAML format
├── internal/
│   ├── domain/                         # Domain Layer (Core Application)
│   │   ├── auth/
│   │   │   ├── dto.go                  # DTOs for registration & login
│   │   │   ├── handler.go              # Handlers for authentication endpoints
│   │   │   └── service.go              # Business logic for authentication
│   │   ├── role/
│   │   │   ├── dto.go                  # DTOs for Role operations
│   │   │   ├── handler.go              # Handlers for Role endpoints
│   │   │   ├── model.go                # Domain models for Role
│   │   │   ├── repository.go           # Repository interfaces (contracts) for Role
│   │   │   └── service.go              # Business logic for Role management
│   │   └── user/
│   │       ├── dto.go                  # DTOs for User operations
│   │       ├── handler.go              # Handlers for User endpoints
│   │       ├── model.go                # Domain models for User
│   │       ├── repository.go           # Repository interfaces (contracts) for User
│   │       └── service.go              # Business logic for User management
│   ├── errors/
│   │   └── errors.go                   # Application custom error definitions
│   └── infrastructure/                 # Infrastructure Layer (Technical Details)
│       ├── cache/
│       │   └── user_cache.go           # User cache repository implementation
│       └── repository/
│           ├── db_models.go            # GORM models for 'users' & 'roles' tables
│           ├── role_repository.go      # Role repository implementation
│           └── user_repository.go      # User repository implementation
├── middlewares/
│   ├── access.go                       # Middleware for Role-Based Access Control (RBAC)
│   └── jwt.go                          # Middleware for JWT token validation
├── routes/
│   └── routes.go                       # API route definitions
├── seeders/
│   ├── role_seeder.go                  # Seeder for default role data
│   └── user.seeder.go                  # Seeder for default user data
├── utils/
│   ├── jwt.go                          # JWT generation & validation utilities
│   ├── password.go                     # Password hashing & comparison utilities
│   └── response.go                     # Standard API response formatting utilities
|
├── .gitignore                          # List of files ignored by Git
├── go.mod                              # Go module declaration
├── go.sum                              # Go module checksums
└── main.go                             # Application entry point

```

*(This structure is based on the uploaded files)*

---

## Installation & Configuration

To run this project locally, follow these steps:

1.  **Clone the Repository**
    ```sh
    git clone https://github.com/nullablenone/user-role-management.git
    cd user-role-management
    ```

2.  **Environment Configuration**
    Copy the `.env.example` file or create a new `.env` file in the project root.
    ```env
    # Database
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=password_anda
    DB_NAME=nama_database
    DB_PORT=5432
    DB_SSLMODE=disable

    # JWT
    SecretKey=kunci_rahasia_jwt_anda

    # Redis
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    REDIS_DB=0
    ```
    *Ensure the variables above match your local configuration.*

3.  **Install Dependencies**
    ```sh
    go mod tidy
    ```

4.  **Run the Application**
    ```sh
    go run main.go
    ```
    The server will start at `http://localhost:8080`.
    > **Note:** The database will be automatically seeded with initial data upon startup.

---

## API Usage

1. **Obtain Token**: Send a `POST` request to `/login` using the default credentials (generated by the seeder):
    - **Email**: `admin@gmail.com`
    - **Password**: `admin@gmail.com`
2. **Use the Token**: Copy the token string from the response. To access protected endpoints, include the `Authorization` header in your request:
    HTTP
    ```
    Authorization: Bearer <your_token>
    ```
3. **Swagger Interface**: Navigate to `http://localhost:8080/swagger/index.html`. Click the **Authorize** button, paste your token, and you can test all endpoints interactively.




