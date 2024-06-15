package handlers

import (
	"html/template"
	"net/http"

	"github.com/foureyez/linkbook/internal/service"
)

type Collections struct {
	templates *template.Template
	svc       service.CollectionService
}

func NewCollectionHandler(templates *template.Template, svc service.CollectionService) Handler {
	return &Collections{
		templates: templates,
		svc:       svc,
	}
}

func (c *Collections) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET /collections", c.getAllCollections)
	mux.HandleFunc("GET /collections/{name}", c.getByName)
	return nil
}

func (c *Collections) getAllCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := c.svc.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := c.templates.ExecuteTemplate(w, "collections", collections); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Collections) getByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	collection, err := c.svc.GetByName(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := c.templates.ExecuteTemplate(w, "collection", collection); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
