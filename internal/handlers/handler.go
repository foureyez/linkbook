package handlers

import (
	"html/template"
	"net/http"
)

type Handler interface {
	RegisterRoutes(mux *http.ServeMux) error
}

func renderErrorTemplate(t *template.Template, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = t.ExecuteTemplate(w, "error.html", nil)
}
