# student-management-backend

This is a simple Student Management REST API built using Golang. The application provides basic CRUD (Create, Read, Update, Delete) operations to manage student records.

It allows users to:
- Add a new student
- Retrieve all students
- Fetch a student by ID
- Update student information
- Delete a student record

## How to run It ?
Create a .env file 

```
cp env.example .env

```
fill the values 

export the env in the terminal 

```
set -a && source .env && set +a

```
run this command in the terminal from the root dir 

```
 go run cmd/app/main.go  
 
```