package main

import (
	"fmt"
	"net/http"
)

// Usuarios "de prueba" (en un proyecto real vendrían de una base de datos)
var users = map[string]string{
	//"admin":  "1234",
	//"user":   "abcd",
	"fehejo": "Fh1969@?",
}

// Handler para /login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Si es POST, procesamos el login
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validar usuario y contraseña
		if pass, ok := users[username]; ok && pass == password {
			// Guardamos sesión en cookie
			http.SetCookie(w, &http.Cookie{
				Name:  "session_user",
				Value: username,
				Path:  "/",
			})
			fmt.Fprintf(w, "✅ Bienvenido, %s!", username)
			return
		}

		// Si no coincide → error
		http.Error(w, "❌ Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	// Si es GET → mostramos formulario
	fmt.Fprint(w, `
		<h1>Login con FEHEJO</h1>
		<form method="POST" action="/login">
			<input name="username" placeholder="Usuario"><br>
			<input type="password" name="password" placeholder="Contraseña"><br>
			<button type="submit">Ingresar</button>
		</form>
	`)
}

// Endpoint privado
func secretHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_user")
	if err != nil || cookie.Value == "" {
		http.Error(w, "⚠️ No estás logueado", http.StatusForbidden)
		return
	}
	fmt.Fprintf(w, "🔒 Acceso permitido. Hola, %s", cookie.Value)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/secret", secretHandler)

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
