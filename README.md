# GoChidori

A Go pet project implementing a clean layered architecture.

## Project Structure

- `cmd/`: Main entry point.
- `internal/`: Private application code.
    - `config/`: Configuration loading (env).
    - `database/`: DB connection and migrations.
    - `handlers/`: HTTP controllers.
    - `models/`: Database entities.
    - `services/`: Business logic layer.
    - `middleware/`: HTTP middlewares (Auth, Logging).
    - `routes/`: API route definitions.

## Getting Started

1. Copy `.env.example` to `.env`.
2. Update the environment variables.
3. Run the project:
   ```bash
   go run cmd/main.go
   ```
