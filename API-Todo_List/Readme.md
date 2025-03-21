# To-Do List API

# 📌 Overview

This project is a simple To-Do List API built using Go, HTML, and JavaScript. It provides basic functionalities to create, list, search, update (mark as completed), and delete tasks. The front end allows users to interact with the API using a simple web interface.

# 🚀 Features

Create new To-Do tasks ✅

List all tasks 📋

Search for a specific task by ID 🔍

Mark a task as completed 🏁

Delete a task ❌

# 📂 Project Structure

![alt text](./Images/image.png)


# 🛠️ Setup and Installation

Prerequisites

Install Go

Install Postman (optional, for testing API endpoints)

Steps to Run


1. Clone the Repository
2. Run the Server
go run main.go
3. Open your browser and go to:
http://localhost:8080

![alt text](./Images/image-1.png)

🔗 API Endpoints

# 1️⃣ Create a New Task

Endpoint: POST /create

Request Body:
```go
{
  "task": "Learn Go",
  "completed": false
}
```

Response:
```go
{
 "id": 1,
  "task": "Learn Go",
  "completed": false
}
```
![alt text](./Images/image-2.png)


# 2️⃣ List All Tasks

Endpoint: GET /list

Response:
```go
[
  {
    "id": 1,
    "task": "Learn Go",
    "completed": false
  }
]
```
![alt text](./Images/image-3.png)


# 3️⃣ Search for a Task by ID

Endpoint: GET /task/{id}

Response:
```go
{
  "id": 1,
  "task": "Learn Go",
  "completed": false
}
```

![alt text](./Images/image-4.png)


# 4️⃣ Mark a Task as Completed
Endpoint: PUT /update/{id}

Response:
```go
{
  "id": 1,
  "task": "Learn Go",
  "completed": true
}
```

![alt text](./Images/image-5.png)

![alt text](./Images/image-6.png)

# 5️⃣ Delete a Task

Endpoint: DELETE /delete/{id}

Response:
```go
{
  "message": "Task deleted successfully"
}
```

![alt text](./Images/image-7.png)

![alt text](./Images/image-8.png)

# 📜 License

This project is open-source and available under the MIT License. Feel free to modify and use it as needed.

💡 Future Enhancements

Add database support (e.g., PostgreSQL, MySQL)

Implement user authentication

Create a more advanced frontend

```go
=============================================================================
DEVELOPER: Aswin KS
DATE: 26-02-2025
ABOUT: Create a TO-Do list using GO, HTML and Javascript
===========================================================================
```