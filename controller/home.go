package controller

import (
	"html/template"
	"net/http"
)

type HomeController struct {
	View *template.Template
}

func (controller *HomeController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	controller.View.ExecuteTemplate(w, "index.html", nil)
}

func Home(view *template.Template) http.Handler {
	return &HomeController{View: view}
}
