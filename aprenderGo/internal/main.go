package main

import (
	"aprenderGo/internal/service"
	"aprenderGo/internal/store"
	"aprenderGo/internal/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func main() {

	//conectar a sqllite
	//db, err := sql.Open("sqlite3", "./books.db")
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//crear la tabla

	q := `
	CREATE TABLE IF NOT EXISTS books (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL, 
	author TEXT NOT NULL
	)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	//inyectar depenencia

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	//configurar las rutas

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookByID)

	fmt.Println("Servidor ejecutandose en http://localhost:8080")
	fmt.Println("API Enpointe:")
	fmt.Println("GET /books  - Obtener todos los libros")
	fmt.Println("POST /books  - Crear un nuevo libro")
	fmt.Println("GET /books/{id}  - Obtener un libro especifico")
	fmt.Println("PUT /books/{id}  - Actualizar un libro")
	fmt.Println("DELETE /books/{id}  - Eliminar un libro")

	//empezar y escuchar el servidor

	log.Fatal(http.ListenAndServe(":8080", nil))
}
