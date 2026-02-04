README.md 
# Notes API with Authentication (Golang)

## 📌 Description
This project is a REST API built with Golang that includes:

- User registration and login
- Password encryption with bcrypt
- Authentication using JWT (JSON Web Tokens)
- Role-based authorization (admin / user)
- CRUD operations for notes
- Data persistence using JSON files

The first registered user is automatically assigned the **admin** role.

---

## 🚀 Technologies Used
- Golang
- net/http
- encoding/json
- JWT (`github.com/golang-jwt/jwt/v5`)
- bcrypt (`golang.org/x/crypto/bcrypt`)

---

## ⚙️ How It Works

- Users register and log in using email and password
- Passwords are stored encrypted
- After login, the API returns a JWT token
- Protected endpoints require a valid token
- Admin users can edit and delete notes
- Regular users can only create and view notes

---

## ▶️ How to Run the Project

### 1️⃣ Requirements
- Go installed (v1.20+ recommended)

### 2️⃣ Run the server
```bash
go run main.go
