---
# Student API with Gin

A RESTful API built with Go, Gin Framework, and SQLite.

## How to Run

1. Install:
 ```bash
   go mod tidy

```

2. Start the server:
```bash
go run main.go

```


*The server will start at `http://localhost:8080`.

---

## Structure

* `models/` - Data structures and input validation rules
* `config/` - Database initialization
* `repositories/` - Database access and SQL queries
* `services/` - Application rules
* `handlers/` - HTTP request/response management

---

## API Endpoints

### 1. Get All Students

* Method: `GET`
* Path: `/students`
* *Response:* `200 OK`

### 2. Get Student by ID

* Method: `GET`
* Path: `/students/:id`
* Response: `200 OK` | `404 Not Found`

### 3. Create Student

* Method: `POST`
* Path: `/students`
* Body (JSON):
* Response: `201 Created` | `400 Bad Request`

### 4. Update Student

* Method: `PUT`
* Path: `/students/:id`
* Body: Same as POST request.
* Response: `200 OK` | `400 Bad Request` | `404 Not Found`

### 5. Delete Student

* Method: `DELETE`
* Path: `/students/:id`
* Response: `204 No Content` | `404 Not Found`

---

## Key Features

* Input Validation:required fields (`id`, `name`) and checks that `gpa` is between `0.00` and `4.00`. Returns `400 Bad Request` if validation fails.
* Error Handling: Standardized JSON error responses.

---