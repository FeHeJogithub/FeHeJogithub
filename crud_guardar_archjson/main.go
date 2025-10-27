package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Edad   int    `json:"edad"`
}

var usuarios []Usuario
var idCounter = 1

const archivo = "usuarios.json"

// 📥 Cargar usuarios desde archivo
func cargarUsuarios() {
	if _, err := os.Stat(archivo); os.IsNotExist(err) {
		return // no existe archivo, iniciamos vacío
	}

	data, err := ioutil.ReadFile(archivo)
	if err != nil {
		fmt.Println("Error leyendo archivo:", err)
		return
	}

	if err := json.Unmarshal(data, &usuarios); err != nil {
		fmt.Println("Error parseando JSON:", err)
		return
	}

	// Actualizar idCounter al mayor ID + 1
	for _, u := range usuarios {
		if u.ID >= idCounter {
			idCounter = u.ID + 1
		}
	}
}

// 📤 Guardar usuarios en archivo
func guardarUsuarios() {
	data, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		fmt.Println("Error convirtiendo JSON:", err)
		return
	}
	ioutil.WriteFile(archivo, data, 0644)
}

// Crear usuario
func crearUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var u Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if u.Nombre == "" || u.Email == "" || u.Edad <= 0 {
		http.Error(w, "Todos los campos son obligatorios y edad > 0", http.StatusBadRequest)
		return
	}

	u.ID = idCounter
	idCounter++
	usuarios = append(usuarios, u)
	guardarUsuarios()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// Listar usuarios
func listarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// Actualizar usuario
func actualizarUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var u Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	for i, usr := range usuarios {
		if usr.ID == id {
			usuarios[i].Nombre = u.Nombre
			usuarios[i].Email = u.Email
			usuarios[i].Edad = u.Edad
			guardarUsuarios()
			json.NewEncoder(w).Encode(usuarios[i])
			return
		}
	}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

// Eliminar usuario
func eliminarUsuario(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, usr := range usuarios {
		if usr.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			guardarUsuarios()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

func main() {
	// Cargar usuarios existentes al iniciar
	cargarUsuarios()

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listarUsuarios(w, r)
		case http.MethodPost:
			crearUsuario(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/usuarios/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			actualizarUsuario(w, r)
		case http.MethodDelete:
			eliminarUsuario(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
