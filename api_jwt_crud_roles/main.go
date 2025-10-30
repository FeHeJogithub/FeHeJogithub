package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("mi_clave_secreta")

type Usuario struct {
	Nombre     string `json:"nombre"`
	Email      string `json:"email"`
	Edad       int    `json:"edad"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"` // 👈 "admin" o "usuario"
}

type Nota struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Contenido string `json:"contenido"`
	Autor     string `json:"autor"`
}

type Claims struct {
	Email string `json:"email"`
	Rol   string `json:"rol"`
	jwt.RegisteredClaims
}

const usuariosFile = "usuarios.json"
const notasFile = "notas.json"

// ---------- Funciones de utilidad ----------
func leerUsuarios() ([]Usuario, error) {
	data, err := os.ReadFile(usuariosFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Usuario{}, nil
		}
		return nil, err
	}
	var usuarios []Usuario
	_ = json.Unmarshal(data, &usuarios)
	return usuarios, nil
}

func guardarUsuarios(usuarios []Usuario) {
	data, _ := json.MarshalIndent(usuarios, "", "  ")
	os.WriteFile(usuariosFile, data, 0644)
}

func leerNotas() ([]Nota, error) {
	data, err := os.ReadFile(notasFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Nota{}, nil
		}
		return nil, err
	}
	var notas []Nota
	_ = json.Unmarshal(data, &notas)
	return notas, nil
}

func guardarNotas(notas []Nota) {
	data, _ := json.MarshalIndent(notas, "", "  ")
	os.WriteFile(notasFile, data, 0644)
}

// ---------- Handlers ----------
func registroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var u Usuario
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if u.Nombre == "" || u.Email == "" || u.Contrasena == "" {
		http.Error(w, "Campos requeridos faltantes", http.StatusBadRequest)
		return
	}

	usuarios, _ := leerUsuarios()
	for _, user := range usuarios {
		if user.Email == u.Email {
			http.Error(w, "Usuario ya registrado", http.StatusBadRequest)
			return
		}
	}

	// Si es el primer usuario registrado, será admin automáticamente
	if len(usuarios) == 0 {
		u.Rol = "admin"
	} else {
		u.Rol = "usuario"
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Contrasena), bcrypt.DefaultCost)
	u.Contrasena = string(hash)
	usuarios = append(usuarios, u)
	guardarUsuarios(usuarios)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario registrado correctamente", "rol": u.Rol})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Email      string `json:"email"`
		Contrasena string `json:"contrasena"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	usuarios, _ := leerUsuarios()
	var user *Usuario
	for _, u := range usuarios {
		if u.Email == req.Email {
			user = &u
			break
		}
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Contrasena), []byte(req.Contrasena)) != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	expTime := time.Now().Add(time.Hour)
	claims := &Claims{
		Email: user.Email,
		Rol:   user.Rol,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtKey)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenStr, "rol": user.Rol})
}

// ---------- Middleware ----------
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token no proporcionado", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
			return
		}

		// Pasamos los claims como contexto
		r.Header.Set("User-Email", claims.Email)
		r.Header.Set("User-Rol", claims.Rol)

		next(w, r)
	}
}

// ---------- CRUD ----------
func notasHandler(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("User-Email")
	notas, _ := leerNotas()

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(notas)

	case http.MethodPost:
		var n Nota
		if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		n.ID = len(notas) + 1
		n.Autor = email
		notas = append(notas, n)
		guardarNotas(notas)
		json.NewEncoder(w).Encode(n)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

var email string

func notaIDHandler(w http.ResponseWriter, r *http.Request) {

	email = r.Header.Get("User-Email")
	rol := r.Header.Get("User-Rol")
	idStr := strings.TrimPrefix(r.URL.Path, "/notas/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	notas, _ := leerNotas()

	switch r.Method {
	case http.MethodPut:
		if rol != "admin" {
			http.Error(w, "Solo los administradores pueden editar notas", http.StatusForbidden)
			return
		}
		var actualizada Nota
		if err := json.NewDecoder(r.Body).Decode(&actualizada); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		for i, n := range notas {
			if n.ID == id {
				notas[i].Titulo = actualizada.Titulo
				notas[i].Contenido = actualizada.Contenido
				guardarNotas(notas)
				json.NewEncoder(w).Encode(notas[i])
				return
			}
		}
		http.Error(w, "Nota no encontrada", http.StatusNotFound)

	case http.MethodDelete:
		if rol != "admin" {
			http.Error(w, "Solo los administradores pueden eliminar notas", http.StatusForbidden)
			return
		}
		for i, n := range notas {
			if n.ID == id {
				notas = append(notas[:i], notas[i+1:]...)
				guardarNotas(notas)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": "Nota eliminada por admin"})
				return
			}
		}
		http.Error(w, "Nota no encontrada", http.StatusNotFound)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	if _, err := os.Stat(usuariosFile); os.IsNotExist(err) {
		os.WriteFile(usuariosFile, []byte("[]"), 0644)
	}
	if _, err := os.Stat(notasFile); os.IsNotExist(err) {
		os.WriteFile(notasFile, []byte("[]"), 0644)
	}

	http.HandleFunc("/registro", registroHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/notas", authMiddleware(notasHandler))
	http.HandleFunc("/notas/", authMiddleware(notaIDHandler))

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
