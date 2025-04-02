# gtodo - Todo List API

`gtodo` is a simple RESTful API built with Go to manage a Todo List. This API allows users to create, read, update, and delete Todo tasks. The application uses a PostgreSQL database to store the todos.

## Features

- **Create Todo**: Add new todos with title, description, status, and due date.
- **Get All Todos**: Retrieve all todos in the system.
- **Get Todo by ID**: Retrieve a specific todo by its ID.
- **Update Todo**: Modify an existing todo.
- **Delete Todo**: Remove a todo by its ID.

## Technologies Used

- **Go**: Backend programming language.
- **PostgreSQL**: Database to store todos.
- **Echo**: Web framework for Go to handle HTTP requests.
- **Gorm**: ORM for Go to interact with the PostgreSQL database.
- **UUID**: For unique identification of todos.
  
## Requirements

- Go (>= 1.16)
- PostgreSQL database
- Git
- Docker (optional, for running the database in a container)

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/gtodo.git
cd gtodo
go mod tidy
```