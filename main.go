package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/whyttea/lenslocked/controllers"
	"github.com/whyttea/lenslocked/views"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template file.", http.StatusInternalServerError)
	}
	tpl.Execute(w, nil)
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {

// 	tplPath := filepath.Join("templates", "home.gohtml")
// 	executeTemplate(w, tplPath)

// IDparam := chi.URLParam(r, "id")
// ctx := r.Context()
// key := ctx.Value("key").(string)

// w.Write([]byte(fmt.Sprintf("New ID is %v, %v", IDparam, key)))
// response := fmt.Sprintf("<h2>%s</h2>", IDparam)
// fmt.Fprint(w, response)

// }

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

//	func pathHandler(w http.ResponseWriter, r *http.Request) {
//		switch r.URL.Path {
//		case "/":
//			homeHandler(w, r)
//		case "/contact":
//			contactHandler(w, r)
//		default:
//			// TODO: add the defualt page
//			http.Error(w, "Page not found", http.StatusNotFound)
//		}
//	}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server at :3000... ")
	http.ListenAndServe("localhost:3000", r)
}
