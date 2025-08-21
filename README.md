# Go Journey

**Go Journey** is a Go-based web application built with the [Fiber](https://gofiber.io/) framework and PostgreSQL as the database.

## Features
- User Management (CRUD)
- Modular Project Structure
- Swagger API Documentation
- Unit Testing with Testify

## Technologies
- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Swagger](https://swagger.io/)
- [Testify](https://github.com/stretchr/testify)

---

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/ganiramadhan/go-journey.git
   cd go-journey
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Setup PostgreSQL database:
   ```sql
   CREATE DATABASE go_journey;
   ```

4. Configure the `.env` file:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=yourpassword
   DB_NAME=go_journey
   APP_PORT=8080
   ```

5. Run the server:
   ```bash
   go run main.go
   ```

6. Access the API at `http://localhost:8080`  
7. Swagger documentation available at `http://localhost:8080/swagger/index.html`

---

## Running Unit Tests

Run all unit tests with:
```bash
go test ./... -v
```

---

## Project Structure

```
go-journey/
│── src/
│   ├── config/         # App and database configuration
│   │   └── config.go
│   ├── controller/     # Request handlers (controllers)
│   │   └── user_controller.go
│   ├── database/       # Database connection & migrations
│   │   ├── migrations/
│   │   └── database.go
│   ├── docs/           # Swagger documentation
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── middleware/     # Application middlewares
│   │   └── auth.go
│   ├── model/          # Data models
│   │   └── user_model.go
│   ├── res/            # Response formatting
│   │   └── user.res.go
│   ├── router/         # Routing definitions
│   ├── service/        # Business logic (services)
│   │   └── user_service.go
│   ├── utils/          # Utility functions
│   ├── validation/     # Validation logic
│   └── test/           # Tests
│       ├── helper/
│       └── unit/
│           └── user_model_test.go
│
│── .env                # Environment variables
│── .env.example        # Example env config
│── .env.production     # Production env config
│── go.mod              # Go module definition
│── main.go             # Application entry point
│── README.md           # Documentation
```

---
