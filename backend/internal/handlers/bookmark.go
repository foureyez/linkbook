package handlers

import (
	"html/template"
	"net/http"

	"github.com/foureyez/linkbook/internal/service"
)

type Bookmark struct {
	templates *template.Template
	svc       service.CollectionService
}

func NewBookmarkHandler(templates *template.Template, svc service.CollectionService) Handler {
	return &Bookmark{
		templates: templates,
		svc:       svc,
	}
}

func (c *Bookmark) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET api/bookmarks", c.getAll)
	mux.HandleFunc("POST api/bookmark", c.create)
	return nil
}

func (c *Bookmark) getAll(w http.ResponseWriter, r *http.Request) {
}

func (c *Bookmark) create(w http.ResponseWriter, r *http.Request) {
}
