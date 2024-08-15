package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my home page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>In case of any issue please contact to: <a href=\"jayantivishnoi@gmail.com\">jayantivishnoi@gmail.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<h1>FAQ Page</h1>
	<ol>
	<li><b>What is duration of the course?</b></li>
	<p>Duration of course id 6 months</p>
	<li><b>Is there are live doubt solving sessions here?</b></li>
	<p>Yes doubt solving session will take place on every weekend</p>
	<li><b>Is the fee refundable?</b></li>
	<p>Yes, it is . You can ask for refund process at any step of course</p</ol>
	`)
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "param")
	fmt.Fprint(w, p)

}

// func pathHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 		return

// 	case "/contact":
// 		contactHandler(w, r)
// 		return
// 	default:
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
		return

	case "/contact":
		contactHandler(w, r)
		return
	case "/faq":
		faqHandler(w, r)
		return
	default:
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintln(w, http.StatusText(http.StatusNotFound))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	// Example handler
	// http.HandleFunc("/", homeHandler)
	// http.HandleFunc("/contact", contactHandler)
	//http.HandleFunc("/", pathHandler)
	// Listen and serve with proper error handling
	//var router Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{param}", paramHandler)
	log.Println("Listening server at :9090")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
