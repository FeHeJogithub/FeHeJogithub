package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)

}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

/*func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	/*m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("TItulo de pagina invalido")
	}
	return m[2], nil

}*/

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[2])
		//return m[2], nil
	}
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	/*title, err := getTitle(w, r)
	if err != nil {
		return
	}*/
	//title := r.URL.Path[len("/view/"):]
	//p, _ := loadPage(title)
	p, err := loadPage(title)
	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplates(w, "view", p)
	/*t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)*/
	//fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}

// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	/*title, err := getTitle(w, r)
	if err != nil {
		return
	}*/

	//title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplates(w, "edit", p)

	/*t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)*/
	//fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}

// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	/*title, err := getTitle(w, r)
	if err != nil {
		return
	}*/

	//title := r.URL.Path[len("/save/"):]

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)

	/*p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}


	renderTemplates(w, "edit", p)*/
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {

	//t, _ := template.ParseFiles("edit.html")
	/*t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

	//err = t.Execute(w, p)
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hola Me encantan los %s", r.URL.Path[1:])
}*/

func main() {

	/*p1 := &Page{Title: "TestPage", Body: []byte("Esta es una pagina de muestra")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))*/

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "Hola Me encantan los %s", "Monos")
	})*/

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
