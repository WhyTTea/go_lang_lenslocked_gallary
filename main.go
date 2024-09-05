package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Println("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template file.", http.StatusInternalServerError)
		return
	}
	viewTpl := views.Template{
		htmlTpl: tpl,
	}
	viewTpl.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)

	// IDparam := chi.URLParam(r, "id")
	// ctx := r.Context()
	// key := ctx.Value("key").(string)

	// w.Write([]byte(fmt.Sprintf("New ID is %v, %v", IDparam, key)))
	// response := fmt.Sprintf("<h2>%s</h2>", IDparam)
	// fmt.Fprint(w, response)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	// alternative to pathing the var with just joining inside of the func call
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

// type Router struct {
// }

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		// TODO: add the defualt page
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// TODO: add the defualt page
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}
// }

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server at :3000... ")
	http.ListenAndServe("localhost:3000", r)
}
