package handlers

import (
	"html/template"
	"net/http"

	"github.com/foureyez/linkbook/internal/logger"
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
	mux.HandleFunc("GET /collections", c.getAll)
	mux.HandleFunc("POST /collections", c.create)
	mux.HandleFunc("GET /collections/{name}", c.getByName)
	mux.HandleFunc("GET /collections/search", c.search)

	return nil
}

func (c *Collections) getAll(w http.ResponseWriter, r *http.Request) {
	collections, err := c.svc.GetAll(r.Context())
	if err != nil {
		logger.Get().Errorf("Unable to get collections, err: %s", err.Error())
		renderErrorTemplate(c.templates, w)
		return
	}

	if err := c.templates.ExecuteTemplate(w, "displayCollections", collections); err != nil {
		logger.Get().Errorf("Unable to serve collections, err: %s", err.Error())
		renderErrorTemplate(c.templates, w)
	}
}

func (c *Collections) create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_, err := c.svc.Create(r.Context(), name)
	if err != nil {
		logger.Get().Errorf("Unable to create collections, err: %s", err.Error())
		renderErrorTemplate(c.templates, w)
		return
	}
	w.Header().Add("HX-Trigger", "refreshCollection")
}

func (c *Collections) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	collections, err := c.svc.SearchByName(r.Context(), query)
	if err != nil {
		logger.Get().Errorf("Unable to get collections, err: %s", err.Error())
		renderErrorTemplate(c.templates, w)
		return
	}

	if err := c.templates.ExecuteTemplate(w, "displayCollections", collections); err != nil {
		logger.Get().Errorf("Unable to serve collections, err: %s", err.Error())
		renderErrorTemplate(c.templates, w)
	}
}

func (c *Collections) getByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	collection, err := c.svc.GetByName(r.Context(), name)
	if err != nil {
		logger.Get().Errorf("Unable to get collection, name: %s,  err: %s", name, err.Error())
		renderErrorTemplate(c.templates, w)
		return
	}

	if err := c.templates.ExecuteTemplate(w, "collection", collection); err != nil {
		logger.Get().Errorf("Unable to serve collection, name: %s,  err: %s", name, err.Error())
		renderErrorTemplate(c.templates, w)
	}
}
