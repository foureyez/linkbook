package service

import persistence "github.com/foureyez/linkbook/internal/peristance"

type Collection struct {
	Name string `json:"name"`
}

func NewCollection(c *persistence.Collection) *Collection {
	return &Collection{
		Name: c.Name,
	}
}
