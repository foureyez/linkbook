package persistence

import "context"

type CollectionStore interface {
	GetAll(ctx context.Context) ([]Collection, error)
	GetByName(ctx context.Context, name string) (*Collection, error)
}
