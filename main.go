package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"webtut/controller"
)

func main() {
	view, err := template.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.Handle("/", controller.Home(view))
	http.Handle("/", router)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":3000", nil)
}
