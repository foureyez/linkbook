package api

import (
	"net/http"

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
	mux.HandleFunc("GET /api/v1/collections", c.getAllCollections)
	mux.HandleFunc("GET /api/v1/collections/{name}", c.getByName)
	return nil
}

func (c *Collections) getAllCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := c.svc.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeJson(w, collections)
}

func (c *Collections) getByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	collection, err := c.svc.GetByName(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeJson(w, collection)
}
