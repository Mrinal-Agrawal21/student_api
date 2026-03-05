# Student API

A simple and robust RESTful API for managing student records, built with Go and SQLite.

## Features
- **Create Student**: Add a new student to the database using `POST /api/v1/students`.
- **Get Student**: Fetch student details by their ID using `GET /api/v1/students/{id}`.
- **Validation**: Strict input validation using `github.com/go-playground/validator`.
- **Database**: Lightweight setup using `go-sqlite3`.
- **Built-in Routing**: Utilizes Go 1.22+'s enhanced `http.ServeMux` for straightforward routing.

## Prerequisites
- **Go**: Version 1.22.0 or higher.
- **SQLite3**: A C compiler requires `CGO_ENABLED=1` for the `go-sqlite3` driver.

## Getting Started

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository_url>
   cd student_api
   ```

2. **Download Dependencies**:
   ```bash
   go mod download
   ```

3. **Configure the Environment**:
   Ensure that a `local.yaml` configuration file exists in `config/local.yaml`:
   ```yaml
   env: "local"
   storage_path: "./storage/storage.db"
   http_server:
     address: "localhost:8080"
     timeout: "4s"
     idle_timeout: "30s"
   ```

4. **Run the Server**:
   ```bash
   go run cmd/student-api/main.go -config config/local.yaml
   ```
   *The server will start locally on `http://localhost:8080` and the SQLite database will be created automatically if it doesn't exist.*

## API Documentation

### 1. Create a Student
- **Endpoint**: `POST /api/v1/students`
- **Headers**: `Content-Type: application/json`
- **Request Body**:
  ```json
  {
      "name": "John Doe",
      "email": "johndoe@example.com",
      "age": 20
  }
  ```
- **Response**: `201 Created`
  ```json
  {
      "id": 1
  }
  ```

### 2. Get a Student by ID
- **Endpoint**: `GET /api/v1/students/{id}`
- **Response**: `200 OK`
  ```json
  {
      "ID": "1",
      "Name": "John Doe",
      "Email": "johndoe@example.com",
      "Age": 20
  }
  ```

## Project Structure
- `cmd/student-api/`: Entrypoint for the application.
- `internal/config/`: Configuration mapping and loader.
- `internal/http/handlers/`: Contains input validation, handler logic, and response formatting.
- `internal/storage/`: Database logic; interactions abstractized via an interface.
- `internal/types/`: Business models and struct definitions.
