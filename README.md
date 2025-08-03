# Task Management API

A RESTful API built with Go Fiber and GORM for managing tasks with user authentication.

## Features

-   User registration and login with JWT authentication
-   CRUD operations for tasks
-   PostgreSQL database with GORM
-   Password hashing with bcrypt
-   UUID for primary keys
-   Middleware for authentication and CORS

## Database Schema

### Users Table

-   `id` (UUID, Primary Key)
-   `username` (String, Unique, Not Null)
-   `password` (String, Not Null, Hashed)
-   `name` (String, Not Null)
-   `created_at` (Timestamp)
-   `updated_at` (Timestamp)

### Tasks Table

-   `id` (UUID, Primary Key)
-   `user_id` (UUID, Foreign Key, Not Null)
-   `todo` (String, Not Null)
-   `start_date` (Timestamp, Not Null)
-   `end_date` (Timestamp, Not Null)
-   `created_at` (Timestamp)
-   `updated_at` (Timestamp)

## API Endpoints

### Authentication

#### Register User

```
POST /api/v1/auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "username": "johndoe",
  "password": "password123"
}
```

#### Login User

```
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "johndoe",
  "password": "password123"
}
```

### Tasks (Protected Routes)

All task endpoints require authentication. Include the JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

#### Get All Tasks

```
GET /api/v1/tasks
```

#### Get Task by ID

```
GET /api/v1/tasks/{id}
```

#### Create Task

```
POST /api/v1/tasks
Content-Type: application/json

{
  "todo": "Complete project documentation",
  "start_date": "2025-08-03T09:00:00Z",
  "end_date": "2025-08-03T17:00:00Z"
}
```

#### Update Task

```
PUT /api/v1/tasks/{id}
Content-Type: application/json

{
  "todo": "Updated task description",
  "start_date": "2025-08-03T10:00:00Z",
  "end_date": "2025-08-03T18:00:00Z"
}
```

#### Delete Task

```
DELETE /api/v1/tasks/{id}
```

## Installation & Setup

1. Clone the repository
2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up PostgreSQL database:

    - Install PostgreSQL on your system
    - Create a database named `tasks_db` (or use your preferred name)
    - Update the environment variables in `.env` file

4. Create a `.env` file (copy from `.env.example`):

    ```bash
    cp .env.example .env
    ```

    Then update the database credentials in the `.env` file.

5. Run the application:

    ```bash
    go run main.go
    ```

6. The server will start on `http://localhost:3000`

## Environment Variables

Create a `.env` file with the following variables:

```
# JWT Secret Key (change this in production)
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Server Configuration
PORT=3000

# PostgreSQL Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password_here
DB_NAME=tasks_db
DB_SSLMODE=disable
```

If no environment variables are provided, the application will use default values (not recommended for production).

## Database

The application uses PostgreSQL database. Make sure you have PostgreSQL installed and running, and create a database before running the application. The tables will be created automatically when you run the application for the first time through GORM's auto-migration feature.

## Testing the API

You can test the API using curl, Postman, or any HTTP client:

### Example: Register and Create Task

1. Register a user:

```bash
curl -X POST http://localhost:3000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "username": "johndoe",
    "password": "password123"
  }'
```

2. Login to get token:

```bash
curl -X POST http://localhost:3000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "password": "password123"
  }'
```

3. Create a task (use the token from login response):

```bash
curl -X POST http://localhost:3000/api/v1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE" \
  -d '{
    "todo": "Complete project documentation",
    "start_date": "2025-08-03T09:00:00Z",
    "end_date": "2025-08-03T17:00:00Z"
  }'
```

## Project Structure

```
.
├── main.go                 # Application entry point
├── database/
│   └── connection.go       # Database configuration
├── models/
│   ├── user.go            # User model
│   └── task.go            # Task model
├── handlers/
│   ├── auth.go            # Authentication handlers
│   └── task.go            # Task CRUD handlers
├── middleware/
│   └── auth.go            # JWT authentication middleware
├── go.mod                 # Go module file
├── go.sum                 # Go dependencies
└── README.md              # This file
```
