package handlers

import (
	"net/http"

	"github.com/foureyez/linkbook/internal/logger"
	"github.com/foureyez/linkbook/internal/service"
)

type Collections struct {
	svc service.CollectionService
}

func NewCollectionHandler(svc service.CollectionService) Handler {
	return &Collections{
		svc: svc,
	}
}

func (c *Collections) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET api/collections", c.getAll)
	mux.HandleFunc("POST api/collections", c.create)
	mux.HandleFunc("GET api/collections/{name}", c.getByName)
	mux.HandleFunc("GET api/collections/search", c.search)

	return nil
}

func (c *Collections) getAll(w http.ResponseWriter, r *http.Request) {
	collections, err := c.svc.GetAll(r.Context())
	if err != nil {
		logger.Get().Errorf("Unable to get collections, err: %s", err.Error())
		return
	}
	encodeJson(w, collections)
}

func (c *Collections) create(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	collection, err := c.svc.Create(r.Context(), name)
	if err != nil {
		logger.Get().Errorf("Unable to create collections, err: %s", err.Error())
		return
	}
	encodeJson(w, collection)
}

func (c *Collections) search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	collections, err := c.svc.SearchByName(r.Context(), query)
	if err != nil {
		logger.Get().Errorf("Unable to get collections, err: %s", err.Error())
		return
	}

	encodeJson(w, collections)
}

func (c *Collections) getByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	collection, err := c.svc.GetByName(r.Context(), name)
	if err != nil {
		logger.Get().Errorf("Unable to get collection, name: %s,  err: %s", name, err.Error())
		return
	}
	encodeJson(w, collection)
}
