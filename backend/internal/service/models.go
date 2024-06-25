package service

import (
	"time"

	persistence "github.com/foureyez/linkbook/internal/peristance"
)

type Collection struct {
	Name         string    `json:"name"`
	ModifiedDate time.Time `json:"modifiedDate"`
}

func NewCollection(c *persistence.Collection) *Collection {
	return &Collection{
		Name:         c.Name,
		ModifiedDate: c.ModifiedDate,
	}
}
