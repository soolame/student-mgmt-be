# student-management-backend

This is a simple **Student Management REST API** built using **Golang**. The application provides basic CRUD (Create, Read, Update, Delete) operations to manage student records.

## Features

It allows users to:

* Add a new student
* Retrieve all students
* Fetch a student by ID
* Update student information
* Delete a student record

---

## How to Run

### 1. Create a `.env` file

```bash
cp env.example .env
```

Fill in the required values in the `.env` file.

---

### 2. Export environment variables

```bash
set -a && source .env && set +a
```

---

### 3. Run the application

From the root directory:

```bash
go run cmd/app/main.go
```

---

## Using Makefile

You can also use the Makefile for common operations:

```bash
make build        # Build the application
make run          # Run the application
make env          # Export environment variables
make migrate-up   # Run database migrations
make migrate-down # Rollback database migrations
```
