# Rock Paper Scissors Web App (Golang)

## 📌 Description
This project is a web application built with Golang that implements the classic **Rock, Paper, Scissors** game.
It uses Go’s standard `net/http` package, custom handlers, and serves static files for the frontend.

---

## 🚀 Technologies Used
- Golang
- net/http
- HTML / CSS
- FileServer for static assets

---

## ⚙️ How It Works
- The application uses an HTTP router (`ServeMux`)
- Routes are handled using custom handler functions
- Static files (CSS, images, JS) are served from the `static` directory
- The server listens on port **8080**

---

## ▶️ How to Run the Project

### 1️⃣ Requirements
- Go installed (v1.20+ recommended)

### 2️⃣ Run the server
```bash
go run main.go
