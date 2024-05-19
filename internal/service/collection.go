package service

import (
	"context"

	"github.com/foureyez/linkbook/internal/logger"
	persistence "github.com/foureyez/linkbook/internal/peristance"
)

type collectionService struct {
	store persistence.CollectionStore
}

func NewCollectionService(store persistence.CollectionStore) CollectionService {
	return &collectionService{
		store: store,
	}
}

func (c *collectionService) GetAll(ctx context.Context) ([]Collection, error) {
	collections, err := c.store.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return transformCollections(collections...), nil
}

func (c *collectionService) GetByName(ctx context.Context, name string) (*Collection, error) {
	collection, err := c.store.GetByName(ctx, name)
	if err != nil {
		logger.Get().Errorf("Unable to get collection, name: %s, err: %s", name, err.Error())
		return nil, err
	}
	return NewCollection(collection), nil
}

func transformCollections(collections ...persistence.Collection) []Collection {
	out := make([]Collection, len(collections))

	for i, c := range collections {
		col := NewCollection(&c)
		out[i] = *col
	}
	return out
}
