# 📝 gtodo - Todo List API

**gtodo** is a simple RESTful API built with Go to manage a Todo List. The API allows users to create, read, update, and delete todo tasks. It uses PostgreSQL to store todos and follows **Clean Architecture** principles for better maintainability and scalability.

---

## 🚀 Features

- ✅ **User Authentication**: JWT-based authentication.
- 🆕 **CRUD Operations For Todo**: Base crud operations for TODO
- 🔗 **Relational Data**: Todo records are related to `User`, `Post`  and `Tag`.

---

## 🧱 Architecture

The project is structured according to **Clean Architecture** principles, making it modular and easy to maintain:

- **`/internal/app`**: Business logic (`usecase`, `service`)
- **`/internal/server`**: Server properties
- **`/cmd/todoapp`**: The entry point of the application
...
---

## 🧰 Technologies Used

| Technology     | Description                              |
|----------------|------------------------------------------|
| **Go**         | Backend programming language             |
| **Echo**       | High-performance web framework           |
| **GORM**       | ORM (Object-Relational Mapping) library  |
| **PostgreSQL** | Database to store todos                  |
| **Docker**     | Containerization                         |
| **Docker Compose** | Used to run the PostgreSQL database in a container |
| **UUID**       | Used for unique identification of todos and users |
| **JWT**        | Token-based authentication for secure access |

---

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/gtodo.git
cd gtodo
go tidy
go run cmd/todoapp/main.go
````

## 🐳 Docker Run Database

First make sure docker is running
```bash
sudo docker-compose -p gtodo up --build
````

## 📌 Contributing
If you'd like to contribute, feel free to submit a pull request or open an issue. Your feedback is appreciated!

# 💻 Developer
Made with ❤️ by [Buğra](https://github.com/bugraozdmr)


