package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

type Player struct {
	Name string
}

var player Player

func Index(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "index.html", nil)
}
func NewGame(w http.ResponseWriter, r *http.Request) {
	restarValue()
	renderTemplate(w, "new-game.html", nil)

	//fmt.Fprintln(w, "Crear nuevo juego")
}
func Game(w http.ResponseWriter, r *http.Request) {
	restarValue()
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}
	//fmt.Println(player.Name)
	renderTemplate(w, "game.html", player)

	//fmt.Fprintln(w, "Juego")
}
func Play(w http.ResponseWriter, r *http.Request) {
	//restarValue()

	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//fmt.Println(result)//resultado en consola
	//fmt.Fprintln(w, "Jugar.html")

}

func About(w http.ResponseWriter, r *http.Request) {
	restarValue()

	renderTemplate(w, "about.html", nil)
	//fmt.Fprintln(w, "Acerca de")
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	//fmt.Fprintln(w, "Hola Mundo con fehejo Pagina de Inicio")
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	//tpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	//tpl, err := template.ParseFiles(`templates\base.html`, `templates\index.html`)
	/*if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
		return
	}*/
	/*data := struct {
		Title   string
		Message string
	}{
		Title:   "Pagina de Inicio con fehejo",
		Message: "!Bienvenido a go con fehejo Piedra , papel o tijera!",
	}*/

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantillas", http.StatusInternalServerError)
		log.Println(err)
		return

	}
}

//reiniciar valores

func restarValue() {
	player.Name = ""

	rps.ComputerScore = 0
	rps.PlayerScore = 0
}
