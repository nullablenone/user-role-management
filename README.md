## Project Structure

```
user-role-management/
├── cmd/
│   └── api/
│       └── main.go                     # App entry point
├── config/
│   ├── database.go                     # Database connection setup
│   └── env.go                          # Environment variable loader
├── internal/
│   └── domain/
│       ├── auth/
│       │   ├── dto.go                  # Data Transfer Objects for auth
│       │   ├── handler.go              # Handler for auth-related endpoints
│       │   └── service.go              # Auth service logic
│       ├── role/
│       │   ├── dto.go                  # DTOs for role operations
│       │   ├── handler.go              # Role handler layer
│       │   ├── model.go                # Role model definition
│       │   ├── repository.go           # Role repository (data access)
│       │   └── service.go              # Role business logic
│       └── user/
│           ├── dto.go                  # DTOs for user operations
│           ├── handler.go              # User handler layer
│           ├── model.go                # User model definition
│           ├── repository.go           # User repository (data access)
│           └── service.go              # User business logic
├── middlewares/
│   ├── access.go                       # Role-based access control middleware
│   └── jwt.go                          # JWT auth middleware
├── routes/
│   └── routes.go                       # All route definitions
├── seeders/
│   ├── role_seeder.go                  # Seeder to populate default role
│   └── user.seeder.go                  # Seeder to populate default users
├── utils/
│   ├── jwt.go                          # JWT creation & parsing
│   ├── password.go                     # Hashing and password comparison utility
│   └── response.go                     # Standardized API response formats
|
├── .env                                # Environment configuration file
├── .gitignore                          # Ignore file
├── go.mod                              # Go module declaration
├── go.sum                              # Go module checksum file
└── README.md                           # Project documentation
```

