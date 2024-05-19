package api

import "net/http"

type Collections struct{}

func (c *Collections) RegisterRoutes(mux *http.ServeMux) error {
	mux.HandleFunc("GET /api/v1/collections", c.getAllCollections)
	return nil
}

func (c *Collections) getAllCollections(w http.ResponseWriter, r *http.Request) {
}
